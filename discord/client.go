package discord

import (
	"github.com/SomethingBot/discordingestor/discord/primitives"
	_ "github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	//apikey is the apikey without a "Bot " prefix
	apikey       string
	intents      primitives.GatewayIntent
	running      bool
	runningLock  sync.Mutex
	eventHandler GatewayEventHandler
}

//New Client using specified apikey without a "Bot " prefix
func New(apikey string, intents primitives.GatewayIntent) *Client {
	return &Client{}
}

func (c *Client) Open() error {
	c.runningLock.Lock()
	defer c.runningLock.Unlock()

	c.running = true
	return nil
}

func (c *Client) Close() error {
	c.runningLock.Lock()
	defer c.runningLock.Unlock()

	c.running = false
	return nil
}

func (c *Client) AddHandlerFunc(i interface{}) error {
	panic("implement me")
}
