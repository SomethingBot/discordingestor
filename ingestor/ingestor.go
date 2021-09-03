package ingestor

import "log"

type Ingestor struct {
	logger         *log.Logger
	discordAPIKey  string
	redisEndpoints []string
}

func New(logger *log.Logger, discordAPIKey string, redisEndpoints []string) Ingestor {
	return Ingestor{
		logger:         logger,
		discordAPIKey:  discordAPIKey,
		redisEndpoints: redisEndpoints,
	}
}

func (ingestor Ingestor) Open() error {
	return nil
}
