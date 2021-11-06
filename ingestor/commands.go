package ingestor

import "github.com/SomethingBot/discordingestor/discordprimatives"

func (ingestor *Ingestor) handleMessages(message discordprimatives.GatewayEventMessageCreate) {
	ingestor.logger.Printf("Saw message (%v)", message.Content)
}
