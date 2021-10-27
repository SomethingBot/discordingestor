package arikawa

import (
	"github.com/SomethingBot/discordingestor/discordprimatives"
	"github.com/SomethingBot/discordingestor/ingestor"
)

//Session implements the DiscordClient and works off an embed Arikawa client
type Session struct {
	//apikey without the "Bot " prefix
	apikey string
}

//New DiscordClient from a Session
func New(apikey string) ingestor.DiscordClient {
	return &Session{apikey: apikey}
}

//Open Arikawa Session
func (arikawaSession *Session) Open() error {
	return nil
}

//Close Arikawa Session
func (arikawaSession *Session) Close() error {
	return nil
}

//AddHandlerFunc to be called on an Event
func (arikawaSession *Session) AddHandlerFunc(handlerFunc func(string)) error {
	return nil
}

//SetIntents to DiscordGateway
func (arikawaSession *Session) SetIntents(discordIntent discordprimatives.GatewayIntent) {

}
