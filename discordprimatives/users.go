package discordprimatives

//ClientStatus struct documented at https://discord.com/developers/docs/topics/gateway#client-status-object
type ClientStatus struct {
	//IsDesktop application session
	IsDesktop bool `json:"desktop"`
	//IsMobile application session
	IsMobile bool `json:"mobile"`
	//IsWeb or Bot application session
	IsWeb bool `json:"web"`
}

//PresenceUpdate struct documented at https://discord.com/developers/docs/topics/gateway#presence-update
type PresenceUpdate struct {
	//User presence is being updated for
	User User `json:"user"`
	//GuildID where PresenceUpdate is for
	GuildID Snowflake `json:"guild_id"`
	//Status that is being updated
	Status PresenceStatus `json:"status"`
	//Activities of User
	Activities []Activity `json:"activities"`
	//ClientStatus of User, platform-dependent
	ClientStatus ClientStatus `json:"client_status"`
}

//User struct from https://discord.com/developers/docs/resources/user#user-object
type User struct {
}
