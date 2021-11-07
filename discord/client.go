package discord

import (
	"github.com/SomethingBot/discordingestor/discord/primatives"
)

type Client struct {
	eventHandler GatewayEventHandler
}

func New() *Client {
	return &Client{}
}

func (c *Client) Open(discordIntent primatives.GatewayIntent) error {
	panic("implement me")
}

func (c *Client) Close() error {
	panic("implement me")
}

func (c *Client) AddHandlerFunc(i interface{}) error {
	panic("implement me")
}
