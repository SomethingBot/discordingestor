package ingestor

import (
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"sync"
)

type ClientState struct {
	runningMutex sync.Mutex
	running      bool
}

type MockDiscordClient struct {
	ClientState
}

func (m *MockDiscordClient) Open() error {
	m.runningMutex.Lock()
	defer func() {
		m.runningMutex.Unlock()
	}()

	if m.running {
		return ErrorSessionAlreadyOpen
	}

	m.running = true
	return nil
}

func (m *MockDiscordClient) Close() error {
	m.runningMutex.Lock()
	defer func() {
		m.runningMutex.Unlock()
	}()

	if !m.running {
		return ErrorSessionAlreadyClosed
	}

	m.running = false
	return nil
}

func (m *MockDiscordClient) AddHandlerFunc(eventType primitives.GatewayEventType, handlerFunc func(event primitives.GatewayEvent)) error {
	return nil
}

func (m *MockDiscordClient) SetIntents(intent primitives.GatewayIntent) {
}

func newMockSessionMaker(apikey string, intents primitives.GatewayIntent) DiscordClient {
	return &MockDiscordClient{}
}
