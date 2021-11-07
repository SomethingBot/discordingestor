package arikawa

import (
	"context"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/primatives"
	"github.com/SomethingBot/discordingestor/ingestor"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

//Arikawa implements the DiscordClient and works off an embed Arikawa client
type Arikawa struct {
	//apikey without the "Bot " prefix
	apikey  string
	session *session.Session
}

//New DiscordClient from a Arikawa
func New(apikey string) ingestor.DiscordClient {
	return &Arikawa{apikey: apikey}
}

//Open Arikawa Arikawa
func (arikawa *Arikawa) Open(discordIntent primatives.GatewayIntent) error {
	var err error
	arikawa.session, err = session.NewWithIntents("Bot "+arikawa.apikey, gateway.Intents(discordIntent))
	if err != nil {
		return fmt.Errorf("Arikawa: failed to open session error (%w)\n", err)
	}

	err = arikawa.session.Open(context.Background())
	if err != nil {
		return err
	}

	return err
}

//Close Arikawa Arikawa
func (arikawa *Arikawa) Close() error {
	err := arikawa.session.Close()
	return err
}

//AddHandlerFunc to be called on an Event
func (arikawa *Arikawa) AddHandlerFunc(f interface{}) error {
	_, err := arikawa.session.AddHandlerCheck(f)
	return fmt.Errorf("Arikawa: failed to add handler error (%w)\n", err)
}
