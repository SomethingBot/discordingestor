package main

import (
	"fmt"
	"github.com/SomethingBot/discordingestor/discord"
	"github.com/SomethingBot/discordingestor/discord/primitives"
	"github.com/SomethingBot/discordingestor/ingestor"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
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

func getDiscordAPIKey() (string, error) {
	var discordAPIKey string
	if discordAPIKeyFile := filepath.Clean(os.Getenv("DISCORD_APIKEYFILE")); discordAPIKeyFile != "" {
		file, err := os.Open(discordAPIKeyFile)
		if err != nil {
			return "", fmt.Errorf("could not open file (%v), error (%v)", discordAPIKeyFile, err)
		}

		fileStat, err := file.Stat()
		if err != nil {
			return "", fmt.Errorf("could not get file (%v) stats (%v)", file.Name(), err)
		}

		if fileStat.Size() == 0 {
			return "", fmt.Errorf("file (%v) size is 0", file.Name())
		}

		if fileStat.Size() > 1<<20 {
			return "", fmt.Errorf("file (%v) size is larger than 1MB", file.Name())
		}

		body, err := io.ReadAll(file)
		if err != nil {
			return "", fmt.Errorf("opened but could not read file (%v), error (%v)", discordAPIKeyFile, err)
		}

		discordAPIKey = string(body)
	}
	return discordAPIKey, nil
}

func main() {
	logger := log.New(os.Stdout, "", log.LUTC|log.Ldate|log.Ltime)
	logger.Printf("Starting discordingestor, commit:%v, tag:%v, Mode:%v", GitCommit, GitTag, Mode)

	redisEndpoints := strings.Split(os.Getenv("REDIS_ENDPOINTS"), ",")

	discordAPIKey, err := getDiscordAPIKey()
	if err != nil {
		logger.Fatalf("error while grabbing Discord API Key (%v)\n", err)
	}

	ingest := ingestor.New(
		logger,
		func(apikey string, intents primitives.GatewayIntent) ingestor.DiscordClient {
			return discord.NewClient(apikey, "", intents, discord.NewEventDistributor())
		},
		ingestor.DiscordConfig{DiscordAPIKey: discordAPIKey},
		ingestor.RedisConfig{RedisEndPoints: redisEndpoints},
	)

	err = ingest.Open()
	if err != nil {
		logger.Printf("Could not open ingestor (%v)", err)
		return
	}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	if signalFromSystem := <-osSignal; signalFromSystem != nil {
		err = ingest.Close()
		if err != nil {
			fmt.Println(err)
		}
		logger.Printf("Stopping discordingestor, commit:%v, tag:%v, Mode:%v, reason: (%v)", GitCommit, GitTag, Mode, signalFromSystem.String())
		return
	}

	logger.Printf("Stopping discordingestor, commit:%v, tag:%v, Mode:%v, reason: (unknown)", GitCommit, GitTag, Mode)
}
