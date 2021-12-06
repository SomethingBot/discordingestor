package discord

import (
	"compress/zlib"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/libinfo"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type discordConfig struct {
	//apikey is the apikey without a "Bot " prefix
	apikey     string
	intents    primitives.GatewayIntent
	gatewayUrl url.URL
}

type discordWebsocket struct {
	conn       *websocket.Conn
	stopChan   chan struct{}
	closeGroup sync.WaitGroup
}

type Client struct {
	discordConfig    discordConfig
	discordWebsocket discordWebsocket
	running          bool
	runningLock      sync.Mutex
	eventDistributor EventDistributor
	logger           *log.Logger
	//todo: ratelimiter
}

//New Client using specified apikey without a "Bot " prefix
func New(apikey string, intents primitives.GatewayIntent, logger *log.Logger) *Client {
	client := &Client{}
	client.logger = logger
	client.discordConfig.apikey = apikey
	client.discordConfig.intents = intents
	client.discordWebsocket.stopChan = make(chan struct{})
	return client
}

func (c *Client) startReadGatewayWorker() {

}

//handleGatewayHandshake separate to prevent having extra checking for events that are not possible
func (c *Client) handleGatewayHandshake() error {
	conn := c.discordWebsocket.conn

	var messageType int
	var err error
	var reader io.Reader
	var event primitives.GEvent //todo: might need to clear seq and other if we get one without and pass it somewhere its used and now has old data from a previous loop
	var decoder *json.Decoder
	for {
		messageType, reader, err = conn.NextReader()
		if err != nil {
			return fmt.Errorf("read error (%w)", err)
		}

		if messageType == websocket.BinaryMessage {
			reader, err = zlib.NewReader(reader)
			if err != nil {
				return err
			}
			//needs a defer or something to prevent leaking zlib reader
		}

		//todo: check for fields sent that are not documented
		decoder = json.NewDecoder(reader)
		err := decoder.Decode(&event)
		if err != nil {
			return fmt.Errorf("error decoding json (%w)", err)
		}

		switch event.Opcode {
		case primitives.GatewayOpcodeHello:
			hello := primitives.GatewayEventHello{} //todo: dont make new variables
			err = json.Unmarshal(event.EventData, &hello)
			if err != nil {
				return fmt.Errorf("error unmarshaling heartbeat (%w)", err)
			}
			//todo: handle heartbeat in goroutine

			//heartbeatAck := `{"op":11}` //todo: this is sent by discord, check if we get one after sending heartbeat, if we dont, reconnect
			//todo: needs to detect if we get an ACK after heartbeating
			//err = conn.WriteMessage(websocket.TextMessage, []byte(heartbeatAck)) //todo: convert to const or something maybe not
			//if err != nil {
			//	return fmt.Errorf("could not write heartbeat ack to websocket")
			//}

			//todo: make actual struct
			identify := `
			{
			  "op": 2,
			  "d": {
				"token": "Bot ` + c.discordConfig.apikey + `",
				"intents": ` + strconv.FormatUint(uint64(c.discordConfig.intents), 10) + `,
				"compress": true,
				"properties": {
				  "$os": "` + runtime.GOOS + `",
				  "$browser": "discordingestor",
				  "$device": "discordingestor"
				}
			  }
			}` //todo: this needs maybe a better name since discordingestor is a pretty non-searchable name

			fmt.Printf("size (%v)\n", len(identify))

			err = conn.WriteMessage(websocket.TextMessage, []byte(identify))
			if err != nil {
				return fmt.Errorf("could not write identify to websocket (%w)", err)
			}
			return nil
		default:
			event = primitives.GetGatewayEventForType(event.EventName)
			err = json.Unmarshal(event.EventData)
		}

		if messageType == websocket.BinaryMessage { //close zlibreader better, maybe dont check messagetype twice (defer?)
			err = reader.(io.ReadCloser).Close()
			if err != nil {
				return fmt.Errorf("error closing zlib readcloser (%w)", err)
			}
		}
	}

	//invalidsession needs a special handle since discord for some reason sends it as a boolean where everything else is an array ???
}

func (c *Client) handleGateway() {

}

func (c *Client) Open() error {
	c.runningLock.Lock()
	defer c.runningLock.Unlock()

	if c.discordConfig.gatewayUrl.String() == "" {
		gURI, err := primitives.GetGatewayURI() //todo: migrate to https://discord.com/developers/docs/topics/gateway#get-gateway-bot
		if err != nil {
			return fmt.Errorf("could not GetGatewayURI (%w)", err)
		}
		c.discordConfig.gatewayUrl = gURI
	}

	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //todo: remove since its for debugging

	gatewayHttpHeader := http.Header{}
	gatewayHttpHeader.Set("User-Agent", libinfo.BotUserAgent)
	gatewayHttpHeader.Set("accept-encoding", "zlib")

	//conn, _, err := dialer.DialContext(context.Background(), c.discordConfig.gatewayUrl.String()+"/?v=9&encoding=json", gatewayHttpHeader)
	//if err != nil {
	//	return err
	//}
	conn, _, err := dialer.DialContext(context.Background(), "wss://localhost:8080"+"/?v=9&encoding=json", gatewayHttpHeader)
	if err != nil {
		return err
	}
	c.discordWebsocket.conn = conn

	defer func() { //todo: make sure to send exit code to discord before dying (if applicable)
		if err != nil {
			err2 := c.closeWebSocket()
			if err2 != nil {
				err = fmt.Errorf("could not close websocket (%v), after error (%w)\n", err2, err)
			}
		}
	}()

	err = c.handleGatewayHandshake()
	if err != nil {
		return fmt.Errorf("discord: gateway handshake error (%w)", err)
	}

	go func(interval int) {
		c.
	}()

	go c.handleGateway()

	c.running = true
	return nil
}

func (c *Client) closeWebSocket() error {
	err := c.discordWebsocket.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	c.runningLock.Lock()
	defer c.runningLock.Unlock()

	if !c.running {
		return fmt.Errorf("discord: client already closed")
	}

	err := c.closeWebSocket()
	if err != nil {
		return err
	}

	c.running = false
	return nil
}

func (c *Client) AddHandlerFunc(i interface{}) error {
	panic("implement me")
}
