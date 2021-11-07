package ingestor

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/primatives"
)

var (
	//ErrorSessionAlreadyOpen returned by a DiscordClient when already open
	ErrorSessionAlreadyOpen = fmt.Errorf("ingestor: session already open")
	//ErrorSessionAlreadyClosed returned by a DiscordClient when already closed
	ErrorSessionAlreadyClosed = fmt.Errorf("ingestor: session already closed")
)

//DiscordClient is the interface called by Ingestor to make Discord API calls and set up the websocket
type DiscordClient interface {
	//Open DiscordClient
	Open(discordIntent primatives.GatewayIntent) error
	//Close DiscordClient
	Close() error
	//AddHandlerFunc that handles an Event
	AddHandlerFunc(interface{}) error
}
