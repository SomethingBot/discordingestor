package discordprimatives

type DiscordGatewayEventType uint32

const (
	//DiscordGatewayEventTypeHello documented at https://discord.com/developers/docs/topics/gateway#hello
	DiscordGatewayEventTypeHello DiscordGatewayEventType = iota
	//DiscordGatewayEventMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
	DiscordGatewayEventMessageCreate
)

type DiscordGatewayEvent interface {
	Type() DiscordGatewayEventType
}

//DiscordGatewayEventHello documented at https://discord.com/developers/docs/topics/gateway#hello
type DiscordGatewayEventHello struct {
}

func (discordGatewayHelloEvent DiscordGatewayEventHello) Type() DiscordGatewayEventType {
	return DiscordGatewayEventTypeHello
}

//EventMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
type EventMessageCreate struct {
	//ID of Message
	ID Snowflake `json:"id"`
	//ChannelID Message was sent in
	ChannelID Snowflake `json:"channel_id"`
	//GuildID Message was sent in
	GuildID Snowflake `json:"guild_id"`
	//Author of Message
	Author User `json:"author"`
	//todo: rest
	//Content of Message
	Content string `json:"content"`
}

func (messageCreateEvent EventMessageCreate) Type() DiscordGatewayEventType {
	return DiscordGatewayEventMessageCreate
}
