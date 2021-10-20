package ingestor

import (
	"fmt"
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
	DiscordAPIKey  string
	discordSession discordSession
}

type ingestorState struct {
	open     bool
	openLock sync.Mutex
}

type Ingestor struct {
	logger       *log.Logger
	sessionMaker func(apikey string) discordSession
	ingestorState
	RedisConfig
	DiscordConfig
}

func New(logger *log.Logger, discordConfig DiscordConfig, redisConfig RedisConfig) *Ingestor {
	return &Ingestor{
		logger:        logger,
		sessionMaker:  newArikawaSession,
		RedisConfig:   redisConfig,
		DiscordConfig: discordConfig,
	}
}

func (ingestor *Ingestor) Open() (err error) {
	ingestor.openLock.Lock()
	defer ingestor.openLock.Unlock()
	if ingestor.open {
		return ErrorInjestorAlreadyOpen
	}

	ingestor.discordSession = ingestor.sessionMaker(strings.TrimSuffix(ingestor.DiscordAPIKey, "\n"))
	if err != nil {
		return err
	}

	ingestor.discordSession.setIntents(discordIntent(gateway.IntentGuildMessages | gateway.IntentGuildInvites | gateway.IntentGuildVoiceStates | gateway.IntentGuilds))

	err = ingestor.discordSession.addHandler(ingestor.handleMessages)
	if err != nil {
		return err
	}

	err = ingestor.discordSession.open()
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

	err = ingestor.discordSession.close()
	if err != nil {
		return err
	}

	ingestor.open = false
	return nil
}
