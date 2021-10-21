package ingestor

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discordprimatives"
	"github.com/diamondburned/arikawa/v3/gateway"
	"log"
	"strings"
	"sync"
)

var (
	ErrorInjestorAlreadyOpen   = fmt.Errorf("ingestor already open")
	ErrorInjestorAlreadyClosed = fmt.Errorf("ingestor already closed")
)

type RedisConfig struct {
	RedisEndPoints []string
}

type DiscordConfig struct {
	DiscordAPIKey string
	DiscordClient DiscordClient
}

type DiscordClientMaker func(apikey string) DiscordClient

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
		return ErrorInjestorAlreadyOpen
	}

	ingestor.DiscordClient = ingestor.discordClientMaker(strings.TrimSuffix(ingestor.DiscordAPIKey, "\n"))
	if err != nil {
		return err
	}

	ingestor.DiscordClient.SetIntents(discordprimatives.Intent(gateway.IntentGuildMessages | gateway.IntentGuildInvites | gateway.IntentGuildVoiceStates | gateway.IntentGuilds))

	err = ingestor.DiscordClient.AddHandler(ingestor.handleMessages)
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
		return ErrorInjestorAlreadyClosed
	}

	err = ingestor.DiscordClient.Close()
	if err != nil {
		return err
	}

	ingestor.open = false
	return nil
}
