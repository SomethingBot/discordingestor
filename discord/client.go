package discord

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/libinfo"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	apikey     string
	gatewayUrl string
	intents    primitives.GatewayIntent

	conn     *websocket.Conn
	sequence syncedCounter

	eDist eDist

	closeErr  error
	closeLock sync.Mutex

	wg         sync.WaitGroup
	senderChan chan senderWork
}

func NewClient(apikey, endpoint string, intents primitives.GatewayIntent, eDist eDist) *Client {
	return &Client{
		apikey:     apikey,
		gatewayUrl: endpoint,
		intents:    intents,
		eDist:      eDist,
		senderChan: make(chan senderWork, 1),
	}
}

type senderWork struct {
	reader   io.Reader
	response chan error
}

func (c *Client) startWebsocketSender() {

}

func (c *Client) sendMessage(reader io.Reader) error {
	response := make(chan error)
	c.senderChan <- senderWork{
		reader:   reader,
		response: nil,
	}
	return <-response
}

func (c *Client) startWebsocketReader() {

}

func (c *Client) handshake() error {
	var err error

	readCloser := io.NopCloser(bytes.NewBuffer(nil)) //todo: replace with a noOpReader, wasting memory for a bit
	defer func() {
		if err != nil {
			err2 := readCloser.Close()
			if err2 != nil {
				err = fmt.Errorf("%w, also could not close readCloser %v", err, err2)
			}
		}
	}()

	var messageType int
	var reader io.Reader
	var gEvent primitives.GEvent
	var decoder *json.Decoder
	for {
		messageType, reader, err = c.conn.NextReader()
		if err != nil {
			return fmt.Errorf("read error (%w)", err)
		}

		if messageType == websocket.BinaryMessage {
			reader, err = zlib.NewReader(reader)
			if err != nil {
				return fmt.Errorf("could not create new zlib reader (%w)", err)
			}
		} else {
			readCloser = io.NopCloser(reader)
		}

		decoder = json.NewDecoder(readCloser)
		err = decoder.Decode(&gEvent)
		if err != nil {
			return fmt.Errorf("could not decode into gEvent (%w)", err)
		}

		c.sequence.set(gEvent.SequenceNumber)

		switch gEvent.Opcode {
		case primitives.GatewayOpcodeHello:
			hello := primitives.GatewayEventHello{}
			err = json.Unmarshal(gEvent.Data, &hello)
			if err != nil {
				return fmt.Errorf("could not unmarshal gEvent.EventDate into a GatewayEventHello (%w)", err)
			}
			c.eDist.FireEvent(hello)
		default:
			var gatewayEvent primitives.GatewayEvent
			gatewayEvent, err = primitives.GetGatewayEventByName(gEvent.Name)
			if err != nil {
				return err
			}
			err = json.Unmarshal(gEvent.Data, &gatewayEvent)
			if err != nil {
				return err
			}
			c.eDist.FireEvent(gatewayEvent)
		}
	}
}

type heartbeat struct {
	primitives.GEvent
}

func (c *Client) startHeartBeatWorker() {
	shutdown := make(chan struct{})
	c.eDist.RegisterHandler(primitives.GatewayEventTypeClientShutdown, func(event primitives.GatewayEvent) {
		shutdown <- struct{}{}
	})
	request := make(chan struct{})
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatRequest, func(event primitives.GatewayEvent) {
		request <- struct{}{}
	})
	ack := make(chan primitives.GatewayEventHeartbeatACK)
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatACK, func(event primitives.GatewayEvent) {
		ack <- event.(primitives.GatewayEventHeartbeatACK) //if this panics, it's an event handler implementation issue, not us
	})
	intervalChange := make(chan int)
	c.eDist.RegisterHandlerOnce(primitives.GatewayEventTypeHello, func(event primitives.GatewayEvent) {
		intervalChange <- event.(primitives.GatewayEventHello).Interval
	})
	go func() {
		interval := <-intervalChange
		close(intervalChange)

		intervalDuration := time.Duration(interval) * time.Millisecond
		timer := time.NewTimer(intervalDuration)
		hasACKed := false
		for {
			select {
			case <-timer.C:
				if !hasACKed {
					c.eDist.FireEvent(primitives.GatewayEventClientShutdown{Err: fmt.Errorf("discord: did not receive an ACK after sending a heartbeat")})
					c.eDist.WaitTilDone()
					return
				}
				//send heartbeat
				timer.Reset(intervalDuration)
			case <-ack:
				hasACKed = true
			case <-request:
				if !timer.Stop() {
					<-timer.C
				}
				hasACKed = false
				//send heartbeat
			case <-shutdown:
				if !timer.Stop() {
					<-timer.C
				}
				c.wg.Done()
			}
		}
	}()
}

func (c *Client) Open() error {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	gatewayHttpHeader := http.Header{}
	gatewayHttpHeader.Set("User-Agent", libinfo.BotUserAgent)
	gatewayHttpHeader.Set("accept-encoding", "zlib")

	var err error
	c.conn, _, err = dialer.DialContext(context.Background(), c.gatewayUrl+"/?v=9&encoding=json", gatewayHttpHeader)
	if err != nil {
		_ = c.conn.Close()
		return err
	}
	defer func() {
		if err != nil {
			err2 := c.conn.Close()
			if err2 != nil {
				err = fmt.Errorf("%w, also could not close c.Conn %v", err, err2)
			}
		}
	}()
	err = c.handshake()
	if err != nil {
		return fmt.Errorf("could not handshake (%w)", err)
	}
	c.startHeartBeatWorker()
	return nil
}

func (c *Client) closeWithError(err error) error { //todo: should return error if closeWithError is called before hand
	c.closeLock.Lock()
	defer c.closeLock.Unlock()

	if c.closeErr != nil {
		return c.closeErr
	}
	c.closeErr = err

	c.eDist.FireEvent(primitives.GatewayEventClientShutdown{Err: fmt.Errorf("discord: did not receive an ACK after sending a heartbeat")})
	c.eDist.WaitTilDone()

	//todo: stop writer
	//todo: stop reader

	return c.closeErr
}

func (c *Client) Close() error {
	return c.closeWithError(nil)
}

func (c *Client) AddHandlerFunc(eventType primitives.GatewayEventType, handlerFunc func(event primitives.GatewayEvent)) error {
	c.eDist.RegisterHandler(eventType, handlerFunc)
	return nil
}
