package ingestor

import "fmt"

var (
	ErrorSessionAlreadyOpen   = fmt.Errorf("ingestor: session already open")
	ErrorSessionAlreadyClosed = fmt.Errorf("ingestor: session already closed")
)

type DiscordClient interface {
	Open() error
	Close() error
	AddHandler(func(string)) error
	SetIntents(DiscordIntent)
}
