package discord

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/primatives"
	"sync"
)

var (
	ErrorEventInvalid                = fmt.Errorf("discord: unknown event")
	ErrorEventHandlerFunctionInvalid = fmt.Errorf("discord: unknown event handler function signature")
)

//GatewayEventHandler is an event handler specifically for primatives.GatewayEvent(s)
type GatewayEventHandler struct {
	handlersLock sync.RWMutex
	handlers     map[primatives.GatewayEventType][]interface{}
}

func newEventHandler() *GatewayEventHandler {
	return &GatewayEventHandler{
		handlers: make(map[primatives.GatewayEventType][]interface{}),
	}
}

//RegisterEventHandlerFunction to be called when an Event is fired
func (eventHandler *GatewayEventHandler) RegisterEventHandlerFunction(handlerFunction interface{}) error {
	eventHandler.handlersLock.Lock() //todo: maybe do this on every line, so we dont lock the entire map while doing a type check
	defer eventHandler.handlersLock.Unlock()
	switch handlerFunction.(type) {
	case func(primatives.GatewayEventHello):
		eventHandler.handlers[primatives.GatewayEventTypeHello] = append(eventHandler.handlers[primatives.GatewayEventTypeHello], handlerFunction)
	default:
		return ErrorEventHandlerFunctionInvalid
	}
	return nil
}

//RemoveEventHandlerFunction and prevent it from being called again
func (eventHandler *GatewayEventHandler) RemoveEventHandlerFunction(handlerFunction interface{}) error {
	eventHandler.handlersLock.Lock()
	defer eventHandler.handlersLock.Unlock()
	switch handlerFunction.(type) {
	case func(primatives.GatewayEventHello):
		for _, handler := range eventHandler.handlers[primatives.GatewayEventTypeHello] {
			if handler == handlerFunction {

			}
		}
	default:
		return ErrorEventHandlerFunctionInvalid
	}
	return nil
}

//FireEvent to registered EventHandlerFunctions; can return a ErrorEventInvalid if EventType is unknown
func (eventHandler *GatewayEventHandler) FireEvent(event primatives.GatewayEvent) error {
	eventHandler.handlersLock.RLock()
	defer eventHandler.handlersLock.RUnlock()
	handlersInterfaces, ok := eventHandler.handlers[event.Type()]
	if !ok {
		return nil
	}

	for _, handler := range handlersInterfaces {
		switch handler.(type) {
		case func(primatives.GatewayEventHello):
			h := handler.(func(primatives.GatewayEventHello)) //todo: maybe return error instead of panic-ing, although if we panic, this is a logic problem
			go h(event.(primatives.GatewayEventHello))
		default:
			return ErrorEventInvalid
		}
	}

	return nil
}
