package ingestor

func (ingestor *Ingestor) handleMessages(message string) {
	//if message.Author.ID == session.State.(*discordgo.State).User.ID {
	//	return
	//}
	ingestor.logger.Printf("Saw message (%v)", message)
}
