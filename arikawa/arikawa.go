package arikawa

import (
	"github.com/SomethingBot/discordingestor/discordprimatives"
	"github.com/SomethingBot/discordingestor/ingestor"
)

type Session struct {
	//apikey without the "Bot " prefix
	apikey string
}

func New(apikey string) ingestor.DiscordClient {
	return &Session{apikey: apikey}
}

func (arikawaSession *Session) Open() error {
	return nil
}

func (arikawaSession *Session) Close() error {
	return nil
}

func (arikawaSession *Session) AddHandler(handlerFunc func(string)) error {
	return nil
}

func (arikawaSession *Session) SetIntents(discordIntent discordprimatives.GatewayIntent) {

}
