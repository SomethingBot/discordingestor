package ingestor

import (
	"log"
	"sync"
	"testing"
)

type mockSession struct {
	runningMutex sync.Mutex
	running      bool
}

func (m *mockSession) open() error {
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

func (m *mockSession) close() error {
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

func (m *mockSession) addHandler(f func(string)) error {
	return nil
}

func (m *mockSession) setIntents(intent discordIntent) {
}

func newMockSessionMaker(apikey string) discordSession {
	return &mockSession{}
}

func TestOpenClose(t *testing.T) {
	t.Parallel()

	ingestor := New(log.Default(), DiscordConfig{}, RedisConfig{})
	ingestor.sessionMaker = newMockSessionMaker

	err := ingestor.Open()
	if err != nil {
		t.Fatalf("ingestor open failed error (%v)\n", err)
	}

	err = ingestor.Open()
	if err != ErrorInjestorAlreadyOpen {
		t.Fatalf("ingestor did not return ErrorIngestorAlreadyOpen when opened twice\n")
	}
	err = nil

	err = ingestor.Close()
	if err != nil {
		t.Fatalf("ingestor close failed error (%v)\n", err)
	}

	err = ingestor.Close()
	if err != ErrorInjestorAlreadyClosed {
		t.Fatalf("ingestor did not return ErrorIngestorAlreadyClosed when closed twice\n")
	}
}
