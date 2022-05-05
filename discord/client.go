package discord

import (
	"compress/zlib"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/libinfo"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"github.com/gorilla/websocket"
	"io"
	mathrand "math/rand"
	"net/http"
	"runtime"
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

	wg          sync.WaitGroup
	senderChan  chan senderWork
	senderClose chan struct{}
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
	data     []byte
	response chan error
}

func (c *Client) startWebsocketWriter() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		var err error
		var work senderWork

		for {
			select {
			case <-c.senderClose:
				return
			case work = <-c.senderChan:
				err = c.conn.WriteMessage(websocket.TextMessage, work.data)
				if err != nil {
					err = c.closeWithError(err)
					if err != nil {
						fmt.Println(fmt.Errorf("failed to closeWithError recieved error (%v)", err))
					}
					work.response <- err
					return
				}
				work.response <- nil
			}
		}
	}()
}

func (c *Client) writeToWebsocket(data []byte) error {
	work := senderWork{
		data:     data,
		response: make(chan error),
	}
	c.senderChan <- work
	return <-work.response
}

//todo: handle resumes

func (c *Client) startWebsocketReader() {

}

func (c *Client) handshake() error {
	messageType, reader, err := c.conn.NextReader()
	if err != nil {
		return fmt.Errorf("could not get next reader (%w)", err)
	}

	var readCloser io.ReadCloser
	if messageType == websocket.BinaryMessage {
		readCloser, err = zlib.NewReader(reader)
		if err != nil {
			return fmt.Errorf("zlib reader could not be created (%w)", err)
		}
		reader = readCloser
		defer func() {
			err2 := readCloser.Close()
			if err2 != nil {
				if err != nil {
					err = fmt.Errorf("could not close zlibReader with error (%v) after error (%w)", err2, err)
				} else {
					err = err2
				}
			}
		}()
	}

	decoder := json.NewDecoder(reader)

	var gEvent primitives.GEvent
	err = decoder.Decode(&gEvent)
	if err != nil {
		return fmt.Errorf("could not decode json gEvent (%w)", err)
	}

	c.sequence.set(gEvent.SequenceNumber)

	//todo: find a way to not have to decode json, store event, then decode json *again*

	var hello primitives.GatewayEventHello
	err = json.Unmarshal(gEvent.Data, &hello)
	if err != nil {
		return fmt.Errorf("could not decode json GatewayEventHello")
	}
	c.eDist.FireEvent(hello)

	var data []byte
	data, err = json.Marshal(primitives.GatewayIdentify{
		Opcode: primitives.GatewayOpcodeIdentify,
		Data: primitives.GatewayIdentifyData{
			Token:   c.apikey,
			Intents: c.intents,
			Properties: primitives.GatewayIdentifyProperties{
				OS:      runtime.GOOS,
				Browser: "discordingestor",
				Device:  "discordingestor",
			},
		},
	})
	if err != nil {
		return err
	}
	err = c.writeToWebsocket(data)
	if err != nil {
		return err
	}
	return nil
}

type heartbeat struct {
	primitives.GEvent
}

func (c *Client) startHeartBeatWorker() error {
	jitter, err := func() (float64, error) {
		b := make([]byte, 8)
		_, err := rand.Read(b)
		if err != nil {
			return 0, err
		}
		/* #nosec G404 */
		return mathrand.New(mathrand.NewSource(int64(binary.BigEndian.Uint64(b)))).Float64(), nil
	}()
	if err != nil {
		return err
	}

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
		intervalChange <- int(float64(event.(primitives.GatewayEventHello).Interval) * jitter)
	})

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		interval := <-intervalChange //todo: handle a resume, which shouldn't require calling this again
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
				err = c.writeToWebsocket([]byte(fmt.Sprintf("{\"op\": 1, \"d\":%v\"}", c.sequence.count())))
				if err != nil {
					_ = c.closeWithError(err)
				}
				timer.Reset(intervalDuration)
			case <-ack:
				hasACKed = true
			case <-request:
				if !timer.Stop() {
					<-timer.C
				}
				hasACKed = false
				err = c.writeToWebsocket([]byte(fmt.Sprintf("{\"op\": 1, \"d\":%v\"}", c.sequence.count())))
				if err != nil {
					_ = c.closeWithError(err)
				}
			case <-shutdown:
				if !timer.Stop() {
					<-timer.C
				}
			}
		}
	}()
	return nil
}

func (c *Client) Open() error {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}

	gatewayHttpHeader := http.Header{}
	gatewayHttpHeader.Set("User-Agent", libinfo.BotUserAgent)
	gatewayHttpHeader.Set("accept-encoding", "zlib")

	var err error
	c.conn, _, err = dialer.DialContext(context.Background(), c.gatewayUrl+"/?v=9&encoding=json", gatewayHttpHeader)
	if err != nil {
		return fmt.Errorf("could not dial (%v) error (%v)", c.gatewayUrl, err)
	}
	defer func() {
		if err != nil {
			err2 := c.conn.Close()
			if err2 != nil {
				err = fmt.Errorf("%w, also could not close c.Conn %v", err, err2)
			}
		}
	}()
	err = c.startHeartBeatWorker()
	if err != nil {
		return err
	}

	c.startWebsocketWriter()

	err = c.handshake()
	if err != nil {
		return fmt.Errorf("could not handshake (%w)", err)
	}

	c.startWebsocketReader()
	return nil
}

func (c *Client) closeWithError(err error) error {
	c.closeLock.Lock()
	defer c.closeLock.Unlock()

	if c.closeErr != nil {
		return c.closeErr
	}
	c.closeErr = err

	c.eDist.FireEvent(primitives.GatewayEventClientShutdown{Err: err})
	c.eDist.WaitTilDone()
	return err
}

func (c *Client) Close() error {
	return c.closeWithError(nil)
}

func (c *Client) AddHandlerFunc(eventType primitives.GatewayEventType, handlerFunc func(event primitives.GatewayEvent)) error {
	c.eDist.RegisterHandler(eventType, handlerFunc)
	return nil
}
