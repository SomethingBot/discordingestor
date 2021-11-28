package main

import (
	"github.com/SomethingBot/discordingestor/discord"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"github.com/SomethingBot/discordingestor/ingestor"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	//GitCommit this package is compiled on
	GitCommit string
	//GitTag this package is compiled on
	GitTag string
	//Mode this package is compiled on
	Mode string
)

func main() {
	logger := log.New(os.Stdout, "", log.LUTC|log.Ldate|log.Ltime)
	logger.Printf("Starting discordingestor, commit:%v, tag:%v, Mode:%v", GitCommit, GitTag, Mode)

	var discordAPIKey string
	if discordAPIKeyFile := os.Getenv("DISCORDAPIKEYFILE"); discordAPIKeyFile != "" {
		file, err := os.Open(discordAPIKeyFile)
		if err != nil {
			logger.Printf("Could not open file (%v), error (%v)", discordAPIKeyFile, err)
			os.Exit(1)
		}

		fileStat, err := file.Stat()
		if err != nil {
			logger.Printf("Could not get file (%v) stats (%v)", file.Name(), err)
		}

		if fileStat.Size() == 0 {
			logger.Printf("DiscordApiKeyFile (%v) size is 0", file.Name())
			os.Exit(1)
		}

		if fileStat.Size() > 1<<20 {
			logger.Printf("Notice: file (%v) size is larger than 1MB", file.Name())
		}

		body, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("Opened, but could not read file (%v), error (%v)", discordAPIKeyFile, err)
			os.Exit(1)
		}

		discordAPIKey = string(body)
	}

	redisEndpoints := strings.Split(os.Getenv("REDISENDPOINTS"), ",")

	ingest := ingestor.New(
		logger,
		func(apikey string, intents primitives.GatewayIntent) ingestor.DiscordClient {
			return discord.New(apikey, intents, logger)
		},
		ingestor.DiscordConfig{DiscordAPIKey: discordAPIKey},
		ingestor.RedisConfig{RedisEndPoints: redisEndpoints},
	)

	err := ingest.Open()
	if err != nil {
		logger.Printf("Could not open ingestor (%v)", err)
		return
	}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	if signalFromSystem := <-osSignal; signalFromSystem != nil {
		logger.Printf("Stopping discordingestor, commit:%v, tag:%v, Mode:%v, reason: (%v)", GitCommit, GitTag, Mode, signalFromSystem.String())
		return
	}

	logger.Printf("Stopping discordingestor, commit:%v, tag:%v, Mode:%v, reason: (unknown)", GitCommit, GitTag, Mode)
}
