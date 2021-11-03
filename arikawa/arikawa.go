package arikawa

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discordprimatives"
	"github.com/SomethingBot/discordingestor/ingestor"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
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
	ariSes, err := session.New("Bot " + arikawaSession.apikey)
	if err != nil {
		return fmt.Errorf("Arikawa: failed to open session error (%w)\n", err) //todo: sentinel error
	}

	//tmp test for echo reply
	_, err = ariSes.AddHandlerCheck(func(msg *gateway.MessageCreateEvent) {

	})
	if err != nil {
		return fmt.Errorf("Arikawa: failed to add handler error (%w)\n", err)
	}

	return nil
}

//Close Arikawa Session
func (arikawaSession *Session) Close() error {
	return nil
}

//AddHandlerFunc to be called on an Event
func (arikawaSession *Session) AddHandlerFunc(function ingestor.EventHandlerFunction) error {
	return nil
}

//SetIntents to DiscordGateway
func (arikawaSession *Session) SetIntents(discordIntent discordprimatives.GatewayIntent) {

}
