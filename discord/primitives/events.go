package primitives

//GatewayEventType documented at https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
type GatewayEventType int

const (
	//GatewayEventTypeHello defines heartbeat interval; documented at https://discord.com/developers/docs/topics/gateway#hello
	GatewayEventTypeHello GatewayEventType = iota
	//GatewayEventTypeReady contains initial state information; documented at https://discord.com/developers/docs/topics/gateway#ready
	GatewayEventTypeReady
	//GatewayEventTypeResumed is the response to a Resume Gateway command
	GatewayEventTypeResumed
	//GatewayEventTypeReconnect is Gateway instructed Client to reconnect and send a Resume
	GatewayEventTypeReconnect
	//GatewayEventTypeInvalidSession is a failure response to a Gateway Identity or a Resume
	GatewayEventTypeInvalidSession
	//GatewayEventTypeChannelCreate is a creation of a Channel
	GatewayEventTypeChannelCreate
	//GatewayEventTypeChannelUpdate is an update of a Channel
	GatewayEventTypeChannelUpdate
	//GatewayEventTypeChannelDelete is a deletion of a Channel
	GatewayEventTypeChannelDelete
	//GatewayEventTypeChannelPinsUpdate is the update of a Channel's pins
	GatewayEventTypeChannelPinsUpdate
	//GatewayEventTypeThreadCreate is the creation of Thread
	GatewayEventTypeThreadCreate
	//GatewayEventTypeThreadUpdate is the update of a Thread
	GatewayEventTypeThreadUpdate
	//GatewayEventTypeThreadDelete is the deletion of a Thread
	GatewayEventTypeThreadDelete
	//GatewayEventTypeThreadListSync is sent when gaining access to a Channel, contains all active Thread(s) in that Channel
	GatewayEventTypeThreadListSync
	//GatewayEventTypeThreadMemberUpdate ThreadMember for bot was updated
	GatewayEventTypeThreadMemberUpdate
	//GatewayEventTypeThreadMembersUpdate multiple ThreadMember(s) were added or removed from a thread
	GatewayEventTypeThreadMembersUpdate
	//GatewayEventTypeGuildCreate lazy-load for unavailable Guild, Guild became available, or User joined a new Guild
	GatewayEventTypeGuildCreate
	//GatewayEventTypeGuildUpdate is the update of a Guild
	GatewayEventTypeGuildUpdate
	//GatewayEventTypeGuildDelete is when a Guild became unavailable, or Bot left/was removed from Guild
	GatewayEventTypeGuildDelete
	//GatewayEventTypeGuildBanAdd is when a User is banned from a Guild
	GatewayEventTypeGuildBanAdd
	//GatewayEventTypeGuildBanRemove is when a User was unbanned from a Guild
	GatewayEventTypeGuildBanRemove
	//GatewayEventTypeGuildEmojisUpdate is a change in Emoji(s) in a Guild
	GatewayEventTypeGuildEmojisUpdate
	//GatewayEventTypeGuildStickersUpdate is a change in Sticker(s) in a Guild
	GatewayEventTypeGuildStickersUpdate
	//GatewayEventTypeGuildIntegrationsUpdate is a change in an Integration(s) in a guild
	GatewayEventTypeGuildIntegrationsUpdate
	//GatewayEventTypeGuildMemberAdd is when a new User joins a Guild
	GatewayEventTypeGuildMemberAdd
	//GatewayEventTypeGuildMemberRemove is when a User leaves or is removed from a Guild
	GatewayEventTypeGuildMemberRemove
	//GatewayEventTypeGuildMemberUpdate is when a GuildMember was updated
	GatewayEventTypeGuildMemberUpdate
	//GatewayEventTypeGuildMembersChunk is a response to a RequestGuildMembers (https://discord.com/developers/docs/topics/gateway#request-guild-members)
	GatewayEventTypeGuildMembersChunk
	//GatewayEventTypeGuildRoleCreate is when a Role is created in a Guild
	GatewayEventTypeGuildRoleCreate
	//GatewayEventTypeGuildRoleUpdate is when a Role is updated in a Guild
	GatewayEventTypeGuildRoleUpdate
	//GatewayEventTypeGuildRoleDelete is when a Role is deleted in a Guild
	GatewayEventTypeGuildRoleDelete
	//GatewayEventTypeGuildIntegrationCreate is when a Guild Integration was created
	GatewayEventTypeGuildIntegrationCreate
	//GatewayEventTypeGuildIntegrationUpdate is when a Guild Integration was updated
	GatewayEventTypeGuildIntegrationUpdate
	//GatewayEventTypeGuildIntegrationDelete is when a Guild Integration was deleted
	GatewayEventTypeGuildIntegrationDelete
	//GatewayEventTypeGuildInteractionCreate is when a User uses an Interaction (like application commands, https://discord.com/developers/docs/interactions/application-commands)
	GatewayEventTypeGuildInteractionCreate
	//GatewayEventTypeInviteCreate is when an Invite to a channel was created
	GatewayEventTypeInviteCreate
	//GatewayEventTypeInviteDelete is when an Invite to a channel was deleted
	GatewayEventTypeInviteDelete
	//GatewayEventTypeMessageCreate documented at https://discord.com/developers/docs/topics/gateway#message-create
	GatewayEventTypeMessageCreate
	//GatewayEventTypeMessageUpdate is when a Message was edited
	GatewayEventTypeMessageUpdate
	//GatewayEventTypeMessageDelete is when a Message was deleted
	GatewayEventTypeMessageDelete
	//GatewayEventTypeMessageDeleteBulk is when multiple Messages were deleted
	GatewayEventTypeMessageDeleteBulk
	//GatewayEventTypeMessageReactionAdd is when a User reacts to a Message
	GatewayEventTypeMessageReactionAdd
	//GatewayEventTypeMessageReactionRemove is when a User removed a reaction from a Message
	GatewayEventTypeMessageReactionRemove
	//GatewayEventTypeMessageReactionRemoveAll is when all reactions were removed from a Message
	GatewayEventTypeMessageReactionRemoveAll
	//GatewayEventTypeMessageReactionRemoveEmoji is when all reactions for a specific Emoji was removed from a Message
	GatewayEventTypeMessageReactionRemoveEmoji
	//GatewayEventTypePresenceUpdate is when a Presence for a User was updated
	GatewayEventTypePresenceUpdate
	//GatewayEventTypeStageInstanceCreate is when a ChannelTypeGuildStageVoice was created in a Guild
	GatewayEventTypeStageInstanceCreate
	//GatewayEventTypeStageInstanceDelete is when a ChannelTypeGuildStageVoice was deleted in a Guild
	GatewayEventTypeStageInstanceDelete
	//GatewayEventTypeStageInstanceUpdate is when a ChannelTypeGuildStageVoice was updated in a Guild
	GatewayEventTypeStageInstanceUpdate
	//GatewayEventTypeTypingStart is when a User has started typing in a Channel
	GatewayEventTypeTypingStart
	//GatewayEventTypeUserUpdate is when a User's properties have been updated
	GatewayEventTypeUserUpdate
	//GatewayEventTypeVoiceStateUpdate is when a User has joined, left, or moved Voice Channel(s); VoiceState
	GatewayEventTypeVoiceStateUpdate
	//GatewayEventTypeVoiceServerUpdate is when a Guild's ChannelTypeGuildVoice has changed Endpoints
	GatewayEventTypeVoiceServerUpdate
	//GatewayEventTypeWebhooksUpdate is when a Guild's Channel's Webhook was created, updated, or deleted
	GatewayEventTypeWebhooksUpdate
)

type GatewayEvent interface {
	//Type of GatewayEvent
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
