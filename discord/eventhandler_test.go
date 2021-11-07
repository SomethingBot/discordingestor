package discord

import (
	"github.com/SomethingBot/discordingestor/discord/primatives"
	"sync"
	"testing"
	"time"
)

type testHandlerUser struct {
	sync.Mutex
	count     int
	firedChan chan bool
}

func (th *testHandlerUser) Handle(helloEvent primatives.GatewayEventHello) {
	th.Lock()
	th.count++
	th.Unlock()
	th.firedChan <- true
}

func TestGatewayEventHandler_FireEvent(t *testing.T) {
	var err error
	eventHandler := newEventHandler()

	thu := &testHandlerUser{firedChan: make(chan bool)}
	err = eventHandler.RegisterEventHandlerFunction(thu.Handle)
	if err != nil {
		t.Fatalf("wanted no error, got error (%v)\n", err)
	}

	handlerUserCount := 100

	for i := 0; i < handlerUserCount; i++ {
		err = eventHandler.FireEvent(primatives.GatewayEventHello{})
		if err != nil {
			t.Fatalf("wanted no error, got err (%v)\n", err)
		}
	}

	timer := time.NewTimer(time.Second)
	for i := 0; i < handlerUserCount; i++ {
		select {
		case fired := <-thu.firedChan:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case <-timer.C:
			t.Fatalf("event fire timed out")
		}
	}

	if thu.count != handlerUserCount {
		t.Fatalf("count is: (%v), wanted count (%v)\n", thu.count, handlerUserCount)
	}

	firedChan := make(chan bool)
	firedChan2 := make(chan bool)
	err = eventHandler.RegisterEventHandlerFunction(func(hello primatives.GatewayEventHello) {
		firedChan <- true
	})
	if err != nil {
		t.Fatalf("wanted no error, got error (%v)\n", err)
	}
	err = eventHandler.RegisterEventHandlerFunction(func(hello primatives.GatewayEventHello) {
		firedChan2 <- true
	})
	if err != nil {
		t.Fatalf("wanted no error, got error (%v)\n", err)
	}

	count := 10

	for i := 0; i < count; i++ {
		err = eventHandler.FireEvent(primatives.GatewayEventHello{})
		if err != nil {
			t.Fatalf("wanted no error, got err (%v)\n", err)
		}
	}

	timer = time.NewTimer(time.Second)
	for i := 0; i < count*2; i++ {
		select {
		case fired := <-firedChan:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case fired := <-firedChan2:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case <-timer.C:
			t.Fatalf("event fire timed out")
		}
	}
}
