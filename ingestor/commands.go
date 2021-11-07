package ingestor

import "github.com/SomethingBot/discordingestor/discord/primatives"

func (ingestor *Ingestor) handleMessages(message primatives.GatewayEventMessageCreate) {
	ingestor.logger.Printf("Saw message (%v)", message.Content)
}
