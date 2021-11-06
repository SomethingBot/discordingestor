package ingestor

import (
	"github.com/SomethingBot/discordingestor/discordprimatives"
	"sync"
)

type ClientState struct {
	runningMutex sync.Mutex
	running      bool
}

type MockDiscordClient struct {
	ClientState
}

func (m *MockDiscordClient) Open(discordIntent discordprimatives.GatewayIntent) error {
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

func (m *MockDiscordClient) AddHandlerFunc(interface{}) error {
	return nil
}

func (m *MockDiscordClient) SetIntents(intent discordprimatives.GatewayIntent) {
}

func newMockSessionMaker(apikey string) DiscordClient {
	return &MockDiscordClient{}
}
