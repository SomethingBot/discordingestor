package discord

import (
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"os"
	"testing"
)

func TestClient_OpenClose(t *testing.T) {
	apikey, err := os.ReadFile("../apikeyfile")
	if err != nil {
		t.Fatalf("error on reading apikeyfile (%v)\n", err)
	}
	client := New(string(apikey), primitives.GatewayIntentGuildMessages)
	err = client.Open()
	if err != nil {
		t.Fatalf("error on open (%v)\n", err)
	}
	err = client.Close()
	if err != nil {
		t.Fatalf("error on close (%v)\n", err)
	}
}
