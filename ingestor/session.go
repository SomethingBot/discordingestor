package ingestor

import "fmt"

var (
	ErrorSessionAlreadyOpen   = fmt.Errorf("ingestor: session already open")
	ErrorSessionAlreadyClosed = fmt.Errorf("ingestor: session already closed")
)

type discordSession interface {
	open() error
	close() error
	addHandler(func(string)) error
	setIntents(discordIntent)
}

type discordIntent uint32
