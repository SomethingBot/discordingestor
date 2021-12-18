package ingestor

import "github.com/SomethingBot/discordingestor/discord/primitives"

func (ingestor *Ingestor) handleMessages(event primitives.GatewayEvent) {
	gc := event.(primitives.GatewayEventGuildCreate)
	ingestor.logger.Printf("Saw guild create (%v)", gc.Name)
}
