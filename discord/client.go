package discord

import (
	"github.com/SomethingBot/discordingestor/discordprimatives"
)

type Client struct {
	eventHandler GatewayEventHandler
}

func New() *Client {
	return &Client{}
}

func (c *Client) Open(discordIntent discordprimatives.GatewayIntent) error {
	panic("implement me")
}

func (c *Client) Close() error {
	panic("implement me")
}

func (c *Client) AddHandlerFunc(i interface{}) error {
	panic("implement me")
}
