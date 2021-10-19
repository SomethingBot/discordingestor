package ingestor

type arikawaSession struct {
	//apikey without the "Bot " prefix
	apikey string
}

func newArikawaSession(apikey string) discordSession {
	return &arikawaSession{apikey: apikey}
}

func (arikawaSession *arikawaSession) open() error {
	panic("implement me")
}

func (arikawaSession *arikawaSession) close() error {
	panic("implement me")
}

func (arikawaSession *arikawaSession) addHandler(handlerFunc func(string)) error {
	panic("implement me")
}

func (arikawaSession *arikawaSession) setIntents(discordIntent discordIntent) {
	panic("implement me")
}
