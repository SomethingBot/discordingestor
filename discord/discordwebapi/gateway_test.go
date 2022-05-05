package discordwebapi

import (
	"os"
	"testing"
)

func TestGetGatewayWebsocketInformation(t *testing.T) {
	t.Parallel()
	//todo: make a short test that doesn't hit server

	if testing.Short() {
		t.Skipf("test short flag set, skipping integration tests")
	}

	apikey := os.Getenv("discordapikey")
	if apikey == "" {
		apikeyBytes, err := os.ReadFile("../../apikeyfile")
		if err != nil {
			t.Fatalf("error on reading apikeyfile (%v)\n", err)
		}
		apikey = string(apikeyBytes)
		if apikey[len(apikey)-1] == '\n' {
			apikey = apikey[:len(apikey)-2]
		}
	}

	_, err := GetGatewayWebsocketInformation("", apikey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetGatewayURI(t *testing.T) {
	t.Parallel()
	//todo: make a short test that doesn't hit server
	uri, err := GetGatewayWebsocketURI("")
	if err != nil {
		t.Fatalf("could not get gateway URI (%v)\n", err)
	}
	if uri.String() == "" {
		t.Fatal("uri returned is empty")
	}
}
