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

func (m *mockSession) Open() error {
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

func (m *mockSession) Close() error {
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

func (m *mockSession) AddHandler(f func(string)) error {
	return nil
}

func (m *mockSession) SetIntents(intent DiscordIntent) {
}

func newMockSessionMaker(apikey string) DiscordClient {
	return &mockSession{}
}

func TestOpenClose(t *testing.T) {
	t.Parallel()

	ingestor := New(log.Default(), newMockSessionMaker, DiscordConfig{}, RedisConfig{})

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
