package ingestor

type discordSession interface {
	open() error
	close() error
	addHandler(func(string)) error
	setIntents(discordIntent)
}

type discordIntent uint32
