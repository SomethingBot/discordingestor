package main

import (
	"github.com/SomethingBot/discordingestor/ingestor"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	GitCommit string
	GitTag    string
	Mode      string
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

	var redisEndpoints []string
	redisEndpointsEnv := os.Getenv("REDISENDPOINTS")
	if !strings.Contains(redisEndpointsEnv, ",") {
		redisEndpoints = []string{redisEndpointsEnv}
	}
	redisEndpoints = strings.Split(redisEndpointsEnv, ",")

	ingest := ingestor.New(logger, ingestor.DiscordConfig{DiscordAPIKey: discordAPIKey}, ingestor.RedisConfig{RedisEndPoints: redisEndpoints})
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
