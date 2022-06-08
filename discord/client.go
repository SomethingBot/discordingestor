package discord

import (
	"compress/zlib"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/libinfo"
	"github.com/SomethingBot/discordingestor/discord/logging"
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

	conn      *websocket.Conn
	sequence  syncedCounter
	writeLock sync.Mutex

	eDist eDist

	closeErr  error
	closeLock sync.Mutex

	logger logging.Logger

	wg sync.WaitGroup

	closeChan chan struct{}
}

func NewClient(apikey, endpoint string, intents primitives.GatewayIntent, eDist eDist, logger logging.Logger) *Client {
	return &Client{
		apikey:     apikey,
		gatewayUrl: endpoint,
		intents:    intents,
		eDist:      eDist,
		logger:     logger,
		closeChan:  make(chan struct{}, 1),
	}
}

type senderWork struct {
	data     []byte
	response chan error
}

func (c *Client) writeToWebsocket(data []byte) error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	c.logger.Log(logging.Debug, fmt.Sprintf("writting message (%v)", string(data)))
	err := c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		go func() {
			_ = c.closeWithError(err)
		}()
	}
	return err
}

//todo: handle resumes

func (c *Client) startWebsocketReader() {
	c.wg.Add(1)
	go func() {
		var err error
		var readCloser io.ReadCloser

		defer func() {
			if readCloser != nil {
				err2 := readCloser.Close()
				if err2 != nil {
					err = fmt.Errorf("could not close readCloser (%v) after error (%w)", err2, err)
				}
			}
			if err != nil {
				c.logger.Log(logging.Error, fmt.Sprintf("WebsocketReader error (%v)", err))
				go func() {
					_ = c.closeWithError(err)
				}()
			}
			c.wg.Done()
		}()

		var messageType int
		var reader io.Reader
		var decoder *json.Decoder
		var gatewayEvent primitives.GatewayEvent
		var gEvent primitives.GEvent

		for {
			select { //todo: this likely won't close until next loop, find way to stop NextReader; currently likely freezes until a heartbeat ack is received
			case _, ok := <-c.closeChan:
				_ = ok
				c.logger.Log(logging.Debug, "Closing WebsocketReader")
				return
			default:
			}

			messageType, reader, err = c.conn.NextReader()
			if err != nil {
				return
			}

			if messageType == websocket.BinaryMessage {
				if readCloser == nil {
					readCloser, err = zlib.NewReader(readCloser)
					if err != nil {
						return
					}
				}

				err = readCloser.(zlib.Resetter).Reset(reader, nil)
				if err != nil {
					return
				}

				reader = readCloser
			}

			decoder = json.NewDecoder(reader)
			err = decoder.Decode(&gEvent)
			if err != nil {
				if errors.Is(err, &(json.UnmarshalTypeError{})) {
					c.logger.Log(logging.Warning, "JSON Data does not unmarshal ("+err.Error()+")")
					err = nil
					continue
				} else {
					return
				}
			}

			c.logger.Log(logging.Debug, "GatewayEvent received ("+gEvent.Name+")")

			c.sequence.set(gEvent.SequenceNumber)

			gatewayEvent, err = primitives.GetGatewayEventByName(gEvent.Name)
			if err != nil {
				c.logger.Log(logging.Warning, "GatewayEvent from Discord not found in primitives.GetGatewayEventByName")
				continue
			}

			err = json.Unmarshal(gEvent.Data, &gatewayEvent)
			if err != nil {
				c.logger.Log(logging.Warning, "JSON Data does not unmarshal ("+err.Error()+") data ("+string(gEvent.Data)+")")
				err = nil
				continue
			}

			c.eDist.FireEvent(gatewayEvent)
		}
	}()
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

var ErrorNoACKAfterHeartbeat = fmt.Errorf("discord: did not receive an ACK after sending a heartbeat")

func (c *Client) generateJitter() float64 { //todo: doesn't need to be on Client struct
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		c.logger.Log(logging.Warning, "Could not access secure rand for generating jitter, falling back to insecure math/rand with current Unix time as seed. Error ("+err.Error()+")")
		/* #nosec G404 */
		return mathrand.New(mathrand.NewSource(time.Now().Unix())).Float64()
	}
	/* #nosec G404 */
	return mathrand.New(mathrand.NewSource(int64(binary.BigEndian.Uint64(b)))).Float64()
}

func (c *Client) startHeartBeatWorker() {
	request := make(chan struct{})
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatRequest, func(event primitives.GatewayEvent) {
		request <- struct{}{}
	})
	ack := make(chan primitives.GatewayEventHeartbeatACK)
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatACK, func(event primitives.GatewayEvent) {
		ack <- event.(primitives.GatewayEventHeartbeatACK)
	})
	intervalChange := make(chan int)
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHello, func(event primitives.GatewayEvent) {
		intervalChange <- int(float64(event.(primitives.GatewayEventHello).Interval) * c.generateJitter())
	})

	c.wg.Add(1)
	go func() {
		interval := <-intervalChange //todo: handle a resume, which shouldn't require calling this again

		intervalDuration := time.Duration(interval) * time.Millisecond
		c.logger.Log(logging.Debug, fmt.Sprintf("starting heartbeat-er with interval (%v)", intervalDuration))

		timer := time.NewTimer(intervalDuration)
		hasACKed := true

		var err error

		defer func() {
			if !timer.Stop() {
				<-timer.C
			}
			if err != nil {
				c.logger.Log(logging.Debug, "heartbeater closed with error ("+err.Error()+")")
				go func() {
					_ = c.closeWithError(err)
				}()
			}
			c.wg.Done()
			c.logger.Log(logging.Debug, "heartbeat worker closed")
		}()

		for {
			select {
			case <-timer.C:
				c.logger.Log(logging.Debug, "Time to heartbeat!")
				if !hasACKed {
					err = ErrorNoACKAfterHeartbeat
					return
				}
				err = c.writeToWebsocket([]byte(fmt.Sprintf("{\"op\": 1, \"d\":%v}", c.sequence.count())))
				if err != nil {
					return
				}
				timer.Reset(intervalDuration)
				c.logger.Log(logging.Debug, "Heartbeat-ed")
			case <-ack:
				hasACKed = true
				c.logger.Log(logging.Debug, "Got heartbeat ack")
			case <-request:
				c.logger.Log(logging.Debug, "Got heartbeat request")
				if !timer.Stop() {
					<-timer.C
				}
				hasACKed = false
				err = c.writeToWebsocket([]byte(fmt.Sprintf("{\"op\": 1, \"d\":%v}", c.sequence.count())))
				if err != nil {
					return
				}
				timer.Reset(intervalDuration)
			case _, _ = <-c.closeChan:
				c.logger.Log(logging.Debug, "Got heartbeat shutdown")
				return
			}
		}
	}()
}

func (c *Client) Open() error {
	c.logger.Log(logging.Info, "Starting Discord Client Library")
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: time.Second, //todo: make configurable
	}

	gatewayHttpHeader := http.Header{}
	gatewayHttpHeader.Set("User-Agent", libinfo.BotUserAgent)
	gatewayHttpHeader.Set("accept-encoding", "zlib") //todo: make configurable

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

	c.eDist.RegisterHandler(primitives.GatewayEventTypeGuildCreate, func(event primitives.GatewayEvent) {
		c.logger.Log(logging.Debug, fmt.Sprintf("guildcreate: (%#v)", event))
	})

	c.startHeartBeatWorker()

	err = c.handshake()
	if err != nil {
		return fmt.Errorf("could not handshake (%w)", err)
	}

	c.startWebsocketReader()
	c.logger.Log(logging.Info, "Client library running")
	return nil
}

func (c *Client) closeWithError(err error) error {
	if err == nil {
		c.logger.Log(logging.Debug, "closeWithError called with nil")
	} else {
		c.logger.Log(logging.Debug, "closeWithError called with ("+err.Error()+")")
	}

	c.closeLock.Lock()
	if c.closeErr != nil {
		c.closeLock.Unlock()
		err2 := c.closeErr
		return err2
	}
	c.closeErr = err
	c.closeLock.Unlock()

	c.eDist.FireEvent(primitives.GatewayEventClientShutdown{Err: err})
	c.eDist.WaitTilDone()

	close(c.closeChan)

	c.wg.Wait()

	err = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	err = c.conn.Close()
	if err != nil {
		return err
	}

	return err
}

func (c *Client) Close() error {
	return c.closeWithError(nil)
}

func (c *Client) AddHandlerFunc(eventType primitives.GatewayEventType, handlerFunc func(event primitives.GatewayEvent)) error {
	c.eDist.RegisterHandler(eventType, handlerFunc)
	return nil
}
