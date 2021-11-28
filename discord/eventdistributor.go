package discord

import (
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"sync"
)

type EventHandler func(event primitives.GatewayEvent)

type handlerEntry struct { //todo: maybe try letting this be locked separately from map containing it
	handlers   []EventHandler
	nilIndexes []int
}

type EventDistributor struct {
	sync.RWMutex
	runningHandlers sync.WaitGroup
	handlerEntries  map[primitives.GatewayEventType]*handlerEntry
}

func NewEventDistributor() *EventDistributor {
	return &EventDistributor{handlerEntries: make(map[primitives.GatewayEventType]*handlerEntry)}
}

func (e *EventDistributor) FireEvent(event primitives.GatewayEvent) {
	e.runningHandlers.Add(1)
	go func() {
		e.RWMutex.RLock()
		handlerEntry, ok := e.handlerEntries[event.Type()]
		if !ok {
			return
		}
		for _, handler := range handlerEntry.handlers {
			if handler != nil {
				e.runningHandlers.Add(1)
				go func(eventHandler EventHandler) {
					eventHandler(event)
					e.runningHandlers.Done()
				}(handler)
			}
		}
		e.RWMutex.RUnlock()
		e.runningHandlers.Done()
	}()
}

type DeregisterFunction func()

type deregisterEntry struct {
	sync.Mutex
	fired            bool
	eventType        primitives.GatewayEventType
	index            int
	eventDistributor *EventDistributor
}

func (d *deregisterEntry) deregisterHandler() {
	d.Lock()
	if d.fired {
		d.Unlock()
		return
	}
	d.fired = true
	d.Unlock()

	d.eventDistributor.RWMutex.Lock()
	entry := d.eventDistributor.handlerEntries[d.eventType]
	entry.handlers[d.index] = nil
	entry.nilIndexes = append(entry.nilIndexes, d.index)
	if len(entry.nilIndexes) == len(entry.handlers) { //todo: this allows us to reset memory usage of slices if we eventually drain all of the handlers, however this is only when *every* handler is removed
		entry.handlers = []EventHandler{}
		entry.nilIndexes = []int{}
	}
	d.eventDistributor.RWMutex.Unlock()
}

func (e *EventDistributor) RegisterHandler(eventType primitives.GatewayEventType, handler EventHandler) DeregisterFunction {
	deregEntry := deregisterEntry{}
	e.RWMutex.Lock()
	entry, ok := e.handlerEntries[eventType]
	if !ok {
		e.handlerEntries[eventType] = &handlerEntry{
			handlers: []EventHandler{handler},
		}
		e.RWMutex.Unlock()
		deregEntry = deregisterEntry{
			eventType:        eventType,
			index:            0,
			eventDistributor: e,
		}
		return deregEntry.deregisterHandler
	}
	if len(entry.nilIndexes) > 0 {
		index := entry.nilIndexes[0]
		entry.handlers[index] = handler
		var slice []int
		slice = append(slice, entry.nilIndexes[1:]...)
		entry.nilIndexes = slice
		e.RWMutex.Unlock()
		deregEntry = deregisterEntry{
			eventType:        eventType,
			index:            index,
			eventDistributor: e,
		}
		return deregEntry.deregisterHandler
	}
	entry.handlers = append(entry.handlers, handler)
	index := len(entry.handlers) - 1

	e.RWMutex.Unlock()
	deregEntry = deregisterEntry{
		eventType:        eventType,
		index:            index,
		eventDistributor: e,
	}
	return deregEntry.deregisterHandler
}

func (e *EventDistributor) WaitTilDone() {
	e.runningHandlers.Wait()
}

type singleFireHandler struct {
	sync.Mutex
	fired              bool
	handler            EventHandler
	deregisterFunction DeregisterFunction
}

func (sfh *singleFireHandler) Fire(event primitives.GatewayEvent) {
	sfh.Lock()
	if sfh.fired {
		sfh.Unlock()
		return
	}
	sfh.handler(event)
	sfh.fired = true
	sfh.deregisterFunction()
	sfh.Unlock()
}

func (e *EventDistributor) RegisterHandlerOnce(eventType primitives.GatewayEventType, handler EventHandler) {
	sfh := singleFireHandler{handler: handler}
	sfh.deregisterFunction = e.RegisterHandler(eventType, sfh.Fire)
}
