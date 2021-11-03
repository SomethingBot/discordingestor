package discordprimatives

type GatewayEventType uint32

const (
	//GatewayEventTypeHello documented at https://discord.com/developers/docs/topics/gateway#hello
	GatewayEventTypeHello GatewayEventType = iota
	//GatewayEventMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
	GatewayEventMessageCreate
)

type GatewayEvent interface {
	Type() GatewayEventType
}

//GatewayEventHello documented at https://discord.com/developers/docs/topics/gateway#hello
type GatewayEventHello struct {
}

func (discordGatewayHelloEvent GatewayEventHello) Type() GatewayEventType {
	return GatewayEventTypeHello
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

func (messageCreateEvent EventMessageCreate) Type() GatewayEventType {
	return GatewayEventMessageCreate
}
