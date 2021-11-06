package discordprimatives

type GatewayEventType int

const (
	//GatewayEventTypeHello documented at https://discord.com/developers/docs/topics/gateway#hello
	GatewayEventTypeHello GatewayEventType = iota
	//GatewayEventTypeMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
	GatewayEventTypeMessageCreate
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

//GatewayEventMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
type GatewayEventMessageCreate struct {
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

func (messageCreateEvent GatewayEventMessageCreate) Type() GatewayEventType {
	return GatewayEventTypeMessageCreate
}
