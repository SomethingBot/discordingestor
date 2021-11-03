package ingestor

import "github.com/SomethingBot/discordingestor/discordprimatives"

func (ingestor *Ingestor) handleMessages(message discordprimatives.EventMessageCreate) {
	ingestor.logger.Printf("Saw message (%v)", message.Content)
}
