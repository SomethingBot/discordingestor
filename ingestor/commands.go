package ingestor

import "github.com/SomethingBot/discordingestor/discord/primitives"

func (ingestor *Ingestor) handleMessages(message primitives.GatewayEventMessageCreate) {
	ingestor.logger.Printf("Saw message (%v)", message.Content)
}
