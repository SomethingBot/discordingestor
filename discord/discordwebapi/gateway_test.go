package discordwebapi

import "testing"

func TestGetGatewayURI(t *testing.T) {
	t.Parallel()
	uri, err := GetGatewayWebsocketURI("")
	if err != nil {
		t.Fatalf("could not get gateway URI (%v)\n", err)
	}
	if uri.String() == "" {
		t.Fatal("uri returned is empty")
	}
}
