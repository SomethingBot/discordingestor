package ingestor

import (
	"log"
	"testing"
)

func TestOpenClose(t *testing.T) {
	t.Parallel()

	ingestor := New(log.Default(), newMockSessionMaker, DiscordConfig{}, RedisConfig{})

	err := ingestor.Open()
	if err != nil {
		t.Fatalf("ingestor open failed error (%v)\n", err)
	}

	err = ingestor.Open()
	if err != ErrorIngestorAlreadyOpen {
		t.Fatalf("ingestor did not return ErrorIngestorAlreadyOpen when opened twice\n")
	}

	err = ingestor.Close()
	if err != nil {
		t.Fatalf("ingestor close failed error (%v)\n", err)
	}

	err = ingestor.Close()
	if err != ErrorIngestorAlreadyClosed {
		t.Fatalf("ingestor did not return ErrorIngestorAlreadyClosed when closed twice\n")
	}
}
