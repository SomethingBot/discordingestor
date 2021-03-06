package ingestor

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"log"
	"strings"
	"sync"
)

var (
	//ErrorIngestorAlreadyOpen is returned by Ingestor when Ingestor.Open() has already been called
	ErrorIngestorAlreadyOpen = fmt.Errorf("ingestor already open")
	//ErrorIngestorAlreadyClosed is returned by Ingestor when Ingestor.Close() has already been called
	ErrorIngestorAlreadyClosed = fmt.Errorf("ingestor already closed")
)

type RedisConfig struct {
	RedisEndPoints []string
}

type DiscordConfig struct {
	DiscordAPIKey string
	DiscordClient DiscordClient
}

type DiscordClientMaker func(apikey string, intents primitives.GatewayIntent) DiscordClient

type ingestorState struct {
	open     bool
	openLock sync.Mutex
}

type Ingestor struct {
	logger             *log.Logger
	discordClientMaker DiscordClientMaker
	ingestorState
	RedisConfig
	DiscordConfig
}

func New(logger *log.Logger, discordClientMaker DiscordClientMaker, discordConfig DiscordConfig, redisConfig RedisConfig) *Ingestor {
	return &Ingestor{
		logger:             logger,
		discordClientMaker: discordClientMaker,
		RedisConfig:        redisConfig,
		DiscordConfig:      discordConfig,
	}
}

func (ingestor *Ingestor) Open() (err error) {
	ingestor.openLock.Lock()
	defer ingestor.openLock.Unlock()
	if ingestor.open {
		return ErrorIngestorAlreadyOpen
	}

	ingestor.DiscordClient = ingestor.discordClientMaker(strings.TrimSuffix(ingestor.DiscordAPIKey, "\n"), primitives.GatewayIntentGuildMessages|primitives.GatewayIntentGuildInvites|primitives.GatewayIntentGuildVoiceStates|primitives.GatewayIntentGuilds)
	if err != nil {
		return err
	}

	err = ingestor.DiscordClient.AddHandlerFunc(primitives.GatewayEventTypeGuildCreate, ingestor.handleMessages)
	if err != nil {
		return err
	}

	err = ingestor.DiscordClient.Open()
	if err != nil {
		return err
	}

	ingestor.open = true
	return nil
}

func (ingestor *Ingestor) Close() (err error) {
	ingestor.openLock.Lock()
	defer ingestor.openLock.Unlock()

	if !ingestor.open {
		return ErrorIngestorAlreadyClosed
	}

	err = ingestor.DiscordClient.Close()
	if err != nil {
		return err
	}

	ingestor.open = false
	return nil
}
