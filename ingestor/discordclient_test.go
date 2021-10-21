package ingestor

import "sync"

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

func (m *MockDiscordClient) AddHandler(f func(string)) error {
	return nil
}

func (m *MockDiscordClient) SetIntents(intent DiscordIntent) {
}

func newMockSessionMaker(apikey string) DiscordClient {
	return &MockDiscordClient{}
}
