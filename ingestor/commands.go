package ingestor

import (
	"github.com/SomethingBot/discordgo"
)

func (ingestor *Ingestor) handleMessages(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.(*discordgo.State).User.ID {
		return
	}
	ingestor.logger.Printf("Saw message (%v)", message.Message)
}
