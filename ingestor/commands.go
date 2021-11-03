package ingestor

import "github.com/SomethingBot/discordingestor/discordprimatives"

func (ingestor *Ingestor) handleMessages(message discordprimatives.DiscordGatewayEvent) {
	event := message.(discordprimatives.EventMessageCreate)
	ingestor.logger.Printf("Saw message (%v)", event.Content)
}
