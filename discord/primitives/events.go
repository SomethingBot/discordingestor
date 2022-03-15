package primitives

import "time"

//go:generate enumer -type GatewayEventType -transform upper -output=events_gatewayevents.go events.go
//GatewayEventType documented at https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
type GatewayEventType int

const (
	//GatewayEventTypeHello defines heartbeat interval; documented at https://discord.com/developers/docs/topics/gateway#hello
	GatewayEventTypeHello GatewayEventType = iota
	//GatewayEventTypeHeartbeatRequest is a request for an immediate heartbeat to be sent to the gateway
	GatewayEventTypeHeartbeatRequest
	//GatewayEventTypeClientShutdown is a shutdown by the DiscordClient Library
	GatewayEventTypeClientShutdown
	//GatewayEventTypeHeartbeatACK is an acknowledgement of a successful heartbeat
	GatewayEventTypeHeartbeatACK
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
	//GatewayEventTypeGuildScheduledEventCreate is when a GuildScheduledEvent is created
	GatewayEventTypeGuildScheduledEventCreate
	//GatewayEventTypeGuildScheduledEventUpdate is when a GuildScheduledEvent is updated
	GatewayEventTypeGuildScheduledEventUpdate
	//GatewayEventTypeGuildScheduledEventDelete is when a GuildScheduledEvent is deleted
	GatewayEventTypeGuildScheduledEventDelete
	//GatewayEventTypeGuildScheduledEventUserAdd is when a GuildScheduledEvent has a User added
	GatewayEventTypeGuildScheduledEventUserAdd
	//GatewayEventTypeGuildScheduledEventUserRemove is when a GuildScheduledEvent has a User removed
	GatewayEventTypeGuildScheduledEventUserRemove
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
	Opcode() GatewayOpcode
}

//GatewayEventHello documented at https://discord.com/developers/docs/topics/gateway#hello
type GatewayEventHello struct {
	Interval int `json:"heartbeat_interval"`
}

func (gatewayHelloEvent GatewayEventHello) Type() GatewayEventType {
	return GatewayEventTypeHello
}

func (gatewayHelloEvent GatewayEventHello) Opcode() GatewayOpcode {
	return GatewayOpcodeHello
}

//UnavailableGuild documented at https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	//ID of UnavailableGuild
	ID Snowflake `json:"id"`
	//IsUnavailable Guild, should always be true
	IsUnavailable bool `json:"unavailable"`
}

//ShardInformation https://discord.com/developers/docs/topics/gateway#sharding
type ShardInformation []int

//ID of Shard
func (s ShardInformation) ID() int {
	return s[0]
}

//Count of Shard(s) Bot should use
func (s ShardInformation) Count() int {
	return s[1]
}

//GatewayEventReady documented at https://discord.com/developers/docs/topics/gateway#ready
type GatewayEventReady struct {
	//Version of Gateway
	Version int `json:"v"`
	//User struct about Bot
	User User `json:"user"`
	//UnavailableGuilds Bot is in
	UnavailableGuilds []UnavailableGuild `json:"unavailable_guilds"`
	//SessionID for resuming connections
	SessionID string `json:"session_id"`
	//ShardInformation associated with SessionID, if sent with Identify
	ShardInformation ShardInformation `json:"shard"`
	//Application of Bot (containing Application.ID and Application.Flags)
	Application Application `json:"application"`
}

//Type of GatewayEventReady
func (gatewayEventReady GatewayEventReady) Type() GatewayEventType {
	return GatewayEventTypeReady
}

//Opcode sent by gateway that contains this event
func (gatewayEventReady GatewayEventReady) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventResumed documented at https://discord.com/developers/docs/topics/gateway#resumed
type GatewayEventResumed struct {
	//Token for Session
	Token string `json:"token"`
	//SessionID for Session
	SessionID string `json:"session_id"`
	//Sequence number last received from Gateway
	Sequence int `json:"seq"`
}

//Type of GatewayEventResumed
func (GatewayEventResumed) Type() GatewayEventType {
	return GatewayEventTypeResumed
}

func (GatewayEventResumed) Opcode() GatewayOpcode {
	return GatewayOpcodeResume
}

//GatewayEventReconnect is documented at https://discord.com/developers/docs/topics/gateway#reconnect
type GatewayEventReconnect struct{}

func (GatewayEventReconnect) Type() GatewayEventType {
	return GatewayEventTypeReconnect
}

func (GatewayEventReconnect) Opcode() GatewayOpcode {
	return GatewayOpcodeReconnect
}

//GatewayEventInvalidSession this struct is probably never used, because discord does not send their data in a sensible way with this specific event, https://discord.com/developers/docs/topics/gateway#invalid-session-example-gateway-invalid-session
type GatewayEventInvalidSession struct{}

func (GatewayEventInvalidSession) Type() GatewayEventType {
	return GatewayEventTypeInvalidSession
}

func (GatewayEventInvalidSession) Opcode() GatewayOpcode {
	return GatewayOpcodeRequestInvalidSession
}

type GatewayEventChannelCreate struct {
	Channel
}

func (GatewayEventChannelCreate) Type() GatewayEventType {
	return GatewayEventTypeChannelCreate
}

func (GatewayEventChannelCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventChannelUpdate documented at https://discord.com/developers/docs/topics/gateway#channel-update; todo: state for "last_message_id" is only tracked when listening for MessageCreate events
type GatewayEventChannelUpdate struct {
	Channel
}

func (GatewayEventChannelUpdate) Type() GatewayEventType {
	return GatewayEventTypeChannelUpdate
}

func (GatewayEventChannelUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

type GatewayEventChannelDelete struct {
	Channel
}

func (GatewayEventChannelDelete) Type() GatewayEventType {
	return GatewayEventTypeChannelDelete
}

func (GatewayEventChannelDelete) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadCreate documented at https://discord.com/developers/docs/topics/gateway#thread-create
//todo: separate event for existing private thread, which has a thread member
type GatewayEventThreadCreate struct {
	Channel
}

func (GatewayEventThreadCreate) Type() GatewayEventType {
	return GatewayEventTypeThreadCreate
}

func (GatewayEventThreadCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadUpdate documented at https://discord.com/developers/docs/topics/gateway#thread-update
//not sent when Channel.LastMessageID is changed, only sent in MessageCreate events
//todo: we need to update the state in the ingestor for MessageCreate events for Channel.LastMessageID
type GatewayEventThreadUpdate struct {
	Channel
}

func (GatewayEventThreadUpdate) Type() GatewayEventType {
	return GatewayEventTypeThreadUpdate
}

func (GatewayEventThreadUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadDelete documented at https://discord.com/developers/docs/topics/gateway#thread-delete
//inner Channel only contains ID, GuildID, ParentID, and Type
type GatewayEventThreadDelete struct {
	Channel
}

func (GatewayEventThreadDelete) Type() GatewayEventType {
	return GatewayEventTypeThreadDelete
}

func (GatewayEventThreadDelete) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadListSync documented at https://discord.com/developers/docs/topics/gateway#thread-list-sync
type GatewayEventThreadListSync struct {
	GuildID    Snowflake      `json:"guild_id"`
	ChannelIDs []Snowflake    `json:"channel_ids"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

func (GatewayEventThreadListSync) Type() GatewayEventType {
	return GatewayEventTypeThreadListSync
}

func (GatewayEventThreadListSync) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadMemberUpdate documented at https://discord.com/developers/docs/topics/gateway#thread-member-update
type GatewayEventThreadMemberUpdate struct {
	ThreadMember
}

func (GatewayEventThreadMemberUpdate) Type() GatewayEventType {
	return GatewayEventTypeThreadMemberUpdate
}

func (GatewayEventThreadMemberUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventThreadMembersUpdate documented at https://discord.com/developers/docs/topics/gateway#thread-members-update
type GatewayEventThreadMembersUpdate struct {
	ID             Snowflake      `json:"id"`
	GuildID        Snowflake      `json:"guild_id"`
	MemberCount    int            `json:"member_count"`
	AddedMembers   []ThreadMember `json:"added_members"`
	RemovedMembers []Snowflake    `json:"removed_member_ids"`
}

func (GatewayEventThreadMembersUpdate) Type() GatewayEventType {
	return GatewayEventTypeThreadMembersUpdate
}

func (GatewayEventThreadMembersUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventChannelPinsUpdate documented at https://discord.com/developers/docs/topics/gateway#channel-pins-update
//This is *NOT* sent when a pinned message is deleted
type GatewayEventChannelPinsUpdate struct {
	GuildID          Snowflake `json:"guild_id"`
	ChannelID        Snowflake `json:"channel_id"`
	LastPinTimestamp time.Time `json:"last_pin_timestamp"`
}

func (GatewayEventChannelPinsUpdate) Type() GatewayEventType {
	return GatewayEventTypeChannelPinsUpdate
}

func (GatewayEventChannelPinsUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildCreate documented at https://discord.com/developers/docs/topics/gateway#guilds
//sent in scenarios: initial connection, after guild is available, Bot joins new Guild
//note: needs GatewayIntentGuildPresences
//Members and Presences in Guild(s) over 75k members contain only the Bot and ChannelTypeGuildVoice User(s)
type GatewayEventGuildCreate struct {
	Guild
}

func (GatewayEventGuildCreate) Type() GatewayEventType {
	return GatewayEventTypeGuildCreate
}

func (GatewayEventGuildCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-update
type GatewayEventGuildUpdate struct {
	Guild
}

func (GatewayEventGuildUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildUpdate
}

func (GatewayEventGuildUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildDelete documented at https://discord.com/developers/docs/topics/gateway#guild-delete
//if unavailable field is not set, user was removed from guild
type GatewayEventGuildDelete struct {
	UnavailableGuild
}

func (GatewayEventGuildDelete) Type() GatewayEventType {
	return GatewayEventTypeGuildDelete
}

func (GatewayEventGuildDelete) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildBanAdd documented at https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GatewayEventGuildBanAdd struct {
	//GuildID User is banned from
	GuildID Snowflake `json:"guild_id"`
	//User that is banned
	User User `json:"user"`
}

func (GatewayEventGuildBanAdd) Type() GatewayEventType {
	return GatewayEventTypeGuildBanAdd
}

func (GatewayEventGuildBanAdd) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildBanRemove documented at https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GatewayEventGuildBanRemove struct {
	//GuildID User is unbanned from
	GuildID Snowflake `json:"guild_id"`
	//User that is unbanned
	User User `json:"user"`
}

func (GatewayEventGuildBanRemove) Type() GatewayEventType {
	return GatewayEventTypeGuildBanRemove
}

func (GatewayEventGuildBanRemove) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildEmojisUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-emojis-update-guild-emojis-update-event-fields
type GatewayEventGuildEmojisUpdate struct {
	//GuildID for Emoji(s) update
	GuildID Snowflake `json:"guild_id"`
	//Emojis of Guild //todo: check if this is actually an array of the changed or just all
	Emojis []Emoji `json:"emojis"`
}

func (GatewayEventGuildEmojisUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildEmojisUpdate
}

func (GatewayEventGuildEmojisUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildStickersUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GatewayEventGuildStickersUpdate struct {
	//GuildID for Sticker(s) update
	GuildID Snowflake `json:"guild_id"`
	//Stickers of Guild //todo: check if this is actually an array of the changed or just all
	Stickers []Sticker `json:"stickers"`
}

func (GatewayEventGuildStickersUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildStickersUpdate
}

func (GatewayEventGuildStickersUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildIntegrationsUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GatewayEventGuildIntegrationsUpdate struct {
	//GuildID for Integration update
	GuildID Snowflake `json:"guild_id"`
}

func (GatewayEventGuildIntegrationsUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildIntegrationsUpdate
}

func (GatewayEventGuildIntegrationsUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildMemberAdd documented at https://discord.com/developers/docs/topics/gateway#guild-member-add
//this requires GatewayIntentGuildMembers
type GatewayEventGuildMemberAdd struct {
	//GuildID of that GuildMember was added to
	GuildID Snowflake `json:"guild_id"`
	GuildMember
}

func (GatewayEventGuildMemberAdd) Type() GatewayEventType {
	return GatewayEventTypeGuildMemberAdd
}

func (GatewayEventGuildMemberAdd) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildMemberRemove documented at https://discord.com/developers/docs/topics/gateway#guild-member-remove
//requires GatewayIntentGuildMembers
type GatewayEventGuildMemberRemove struct {
	//GuildID of the User removed
	GuildID Snowflake `json:"guild_id"`
	//User removed from Guild
	User User `json:"user"`
}

func (GatewayEventGuildMemberRemove) Type() GatewayEventType {
	return GatewayEventTypeGuildMemberRemove
}

func (GatewayEventGuildMemberRemove) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildMemberUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-member-update
//requires GatewayIntentGuildMembers
type GatewayEventGuildMemberUpdate struct {
	//GuildID of the User updated
	GuildID Snowflake `json:"guild_id"`
	//Roles of User
	Roles []Role `json:"roles"`
	//User updated
	User User `json:"user"`
	//Nickname of User
	Nickname string `json:"nick"`
	//AvatarHash of User
	AvatarHash ImageHash `json:"avatar_hash"`
	//JoinedAt time to Guild
	JoinedAt time.Time `json:"joined_at"`
	//PremiumSince Time in Guild
	PremiumSince time.Time `json:"premium_since"`
	//IsDeafened in Voice Channels
	IsDeafened bool `json:"deaf"`
	//IsMuted in Voice Channels
	IsMuted bool `json:"is_muted"`
	//IsPending membership screening
	IsPending bool `json:"is_pending"`
	//CommunicationDisabledUntil time.Time that the GuildMember will be able to communication is enabled again
	CommunicationDisabledUntil time.Time `json:"communication_disabled_until"`
}

func (GatewayEventGuildMemberUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildMemberUpdate
}

func (GatewayEventGuildMemberUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildMembersChunk documented at https://discord.com/developers/docs/topics/gateway#guild-members-chunk
//send in response to GuildRequestMembers
type GatewayEventGuildMembersChunk struct {
	//GuildID of the User updated
	GuildID Snowflake `json:"guild_id"`
	//Members of Chunk
	Members []GuildMember `json:"members"`
	//Index of Chunk
	Index int `json:"index"`
	//Count of Chunks expected
	Count int `json:"count"`
	//NotFoundIDs is the invalid ID sent to GuildRequestMembers
	//todo: check if this is a string
	NotFoundIDs []string `json:"not_found"`
	//Presences if true passed to GuildRequestMembers
	Presences []PresenceUpdate `json:"presences"`
	//Nonce used in GuildRequestMembers
	Nonce string `json:"nonce"`
}

func (GatewayEventGuildMembersChunk) Type() GatewayEventType {
	return GatewayEventTypeGuildMembersChunk
}

func (GatewayEventGuildMembersChunk) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildRoleCreate documented at https://discord.com/developers/docs/topics/gateway#guild-role-create
type GatewayEventGuildRoleCreate struct {
	//GuildID of the Role
	GuildID Snowflake `json:"guild_id"`
	//Role created
	Role Role `json:"role"`
}

func (GatewayEventGuildRoleCreate) Type() GatewayEventType {
	return GatewayEventTypeGuildRoleCreate
}

func (GatewayEventGuildRoleCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildRoleUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-role-update
type GatewayEventGuildRoleUpdate struct {
	//GuildID of the Role
	GuildID Snowflake `json:"guild_id"`
	//Role created
	Role Role `json:"role"`
}

func (GatewayEventGuildRoleUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildRoleUpdate
}

func (GatewayEventGuildRoleUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildRoleDelete documented at https://discord.com/developers/docs/topics/gateway#guild-role-delete
type GatewayEventGuildRoleDelete struct {
	//GuildID of the Role
	GuildID Snowflake `json:"guild_id"`
	//Role created
	RoleID Snowflake `json:"role_id"`
}

func (GatewayEventGuildRoleDelete) Type() GatewayEventType {
	return GatewayEventTypeGuildRoleDelete
}

func (GatewayEventGuildRoleDelete) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildScheduledEventCreate documented at https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-create
type GatewayEventGuildScheduledEventCreate struct {
	GuildScheduledEvent
}

func (GatewayEventGuildScheduledEventCreate) Type() GatewayEventType {
	return GatewayEventTypeGuildScheduledEventCreate
}

func (GatewayEventGuildScheduledEventCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildScheduledEventUpdate documented at https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-update
type GatewayEventGuildScheduledEventUpdate struct {
	GuildScheduledEvent
}

func (GatewayEventGuildScheduledEventUpdate) Type() GatewayEventType {
	return GatewayEventTypeGuildScheduledEventUpdate
}

func (GatewayEventGuildScheduledEventUpdate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildScheduledEventDelete documented at https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-delete
type GatewayEventGuildScheduledEventDelete struct {
	GuildScheduledEvent
}

func (GatewayEventGuildScheduledEventDelete) Type() GatewayEventType {
	return GatewayEventTypeGuildScheduledEventDelete
}

func (GatewayEventGuildScheduledEventDelete) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildScheduledEventUserAdd documented at https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-add
type GatewayEventGuildScheduledEventUserAdd struct {
	//GuildScheduledEventID is the ID of the GuildScheduledEvent
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	//UserID of the User added to the GuildScheduledEvent
	UserID Snowflake `json:"user_id"`
	//GuildID where GuildScheduledEvent is taking place
	GuildID Snowflake `json:"guild_id"`
}

func (GatewayEventGuildScheduledEventUserAdd) Type() GatewayEventType {
	return GatewayEventTypeGuildScheduledEventUserAdd
}

func (GatewayEventGuildScheduledEventUserAdd) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventGuildScheduledEventUserRemove documented at https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-remove
type GatewayEventGuildScheduledEventUserRemove struct {
	//GuildScheduledEventID is the ID of the GuildScheduledEvent
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	//UserID of the User removed from the GuildScheduledEvent
	UserID Snowflake `json:"user_id"`
	//GuildID where GuildScheduledEvent is taking place
	GuildID Snowflake `json:"guild_id"`
}

func (GatewayEventGuildScheduledEventUserRemove) Type() GatewayEventType {
	return GatewayEventTypeGuildScheduledEventUserRemove
}

func (GatewayEventGuildScheduledEventUserRemove) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventIntegrationCreate documented at https://discord.com/developers/docs/topics/gateway#integration-create
type GatewayEventIntegrationCreate struct {
	//GuildID is the ID of the Guild the Integration is created in
	GuildID Snowflake `json:"guild_id"`
}

func (GatewayEventIntegrationCreate) Type() GatewayEventType {
	return GatewayEventTypeGuildIntegrationCreate
}

func (GatewayEventIntegrationCreate) Opcode() GatewayOpcode {
	return GatewayOpcodeDispatch
}

//GatewayEventClientShutdown is the event thrown when the Client Library is shutdown
type GatewayEventClientShutdown struct {
	//err that caused shutdown
	err error
}

func (GatewayEventClientShutdown) Type() GatewayEventType {
	return GatewayEventTypeClientShutdown
}

func (GatewayEventClientShutdown) Opcode() GatewayOpcode {
	return GatewayOpcodeNil
}

//Error that caused the shutdown, nil if no error and standard shutdown
func (g GatewayEventClientShutdown) Error() error {
	return g.err
}
