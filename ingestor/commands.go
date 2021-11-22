package ingestor

import "github.com/SomethingBot/discordingestor/discord/primitives"

func (ingestor *Ingestor) handleMessages(gc primitives.GatewayEventGuildCreate) {
	ingestor.logger.Printf("Saw guild create (%v)", gc.Name)
}
