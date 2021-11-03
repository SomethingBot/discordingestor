package ingestor

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discordprimatives"
)

var (
	//ErrorSessionAlreadyOpen returned by a DiscordClient when already open
	ErrorSessionAlreadyOpen = fmt.Errorf("ingestor: session already open")
	//ErrorSessionAlreadyClosed returned by a DiscordClient when already closed
	ErrorSessionAlreadyClosed = fmt.Errorf("ingestor: session already closed")
)

type EventHandlerFunction func(interface{})

//DiscordClient is the interface called by Ingestor to make Discord API calls and set up the websocket
type DiscordClient interface {
	//Open DiscordClient
	Open() error
	//Close DiscordClient
	Close() error
	//AddHandlerFunc that handles an Event
	AddHandlerFunc(EventHandlerFunction) error
	//SetIntents to the Gateway; call before Open
	SetIntents(discordprimatives.GatewayIntent)
}
