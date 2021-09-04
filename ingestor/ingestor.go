package ingestor

import (
	"fmt"
	"github.com/SomethingBot/discordgo"
	"log"
	"strings"
	"sync"
)

var (
	ErrorAlreadyOpen   = fmt.Errorf("ingestor already open")
	ErrorAlreadyClosed = fmt.Errorf("ingestor already closed")
)

type RedisConfig struct {
	redisEndPoints []string
}

type DiscordConfig struct {
	discordAPIKey    string
	discordGoSession *discordgo.Session
}

type Ingestor struct {
	logger   *log.Logger
	open     bool
	openLock sync.Mutex
	RedisConfig
	DiscordConfig
}

func New(logger *log.Logger, discordAPIKey string, redisEndpoints []string) (ingestor Ingestor) {
	ingestor.logger = logger
	ingestor.discordAPIKey = discordAPIKey
	ingestor.redisEndPoints = redisEndpoints
	return
}

func (ingestor *Ingestor) Open() (err error) {
	ingestor.openLock.Lock()
	defer ingestor.openLock.Unlock()
	if ingestor.open {
		return ErrorAlreadyOpen
	}
	ingestor.open = true

	ingestor.discordGoSession, err = discordgo.New("Bot " + strings.TrimSuffix(ingestor.discordAPIKey, "\n"))
	if err != nil {
		return err
	}

	ingestor.discordGoSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuildInvites | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuilds)

	ingestor.discordGoSession.AddHandler(ingestor.handleMessages)

	err = ingestor.discordGoSession.Open()
	if err != nil {
		return err
	}

	return nil
}

func (ingestor *Ingestor) Close() (err error) {
	ingestor.openLock.Lock()
	defer ingestor.openLock.Unlock()
	if !ingestor.open {
		return ErrorAlreadyClosed
	}
	ingestor.open = false

	err = ingestor.discordGoSession.Close()
	if err != nil {
		return err
	}

	return nil
}
