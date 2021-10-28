package discordprimatives

//GatewayIntent from https://discord.com/developers/docs/topics/gateway#gateway-intents
type GatewayIntent uint16

const (
	//GatewayIntentNil is when no GatewayIntent is set
	GatewayIntentNil GatewayIntent = 0
	//GatewayIntentGuilds contains events:
	//- GuildCreate
	//- GuildUpdate
	//- GuildDelete
	//- GuildRoleCreate
	//- GuildRoleUpdate
	//- GuildRoleDelete
	//- ChannelCreate
	//- ChannelUpdate
	//- ChannelDelete
	//- ChannelPinsUpdate
	//- ThreadCreate
	//- ThreadUpdate
	//- ThreadDelete
	//- ThreadListSync
	//- ThreadMemberUpdate
	//- ThreadMembersUpdate; data is different depending on intents used
	//- StageInstanceCreate
	//- StageInstanceUpdate
	//- StageInstanceDelete
	GatewayIntentGuilds GatewayIntent = 1 << (iota - 1)
	//GatewayIntentGuildMembers contains events:
	//- GuildMemberAdd
	//- GuildMemberUpdate
	//- GuildMemberRemove
	//- ThreadMembersUpdate *
	GatewayIntentGuildMembers
	//GatewayIntentGuildBans contains events:
	//- GuildBanAdd
	//- GuildBanRemove
	GatewayIntentGuildBans
	//GatewayIntentGuildEmojisAndStickers contains events:
	//- GuildEmojisUpdate
	//- GuildStickersUpdate
	GatewayIntentGuildEmojisAndStickers
	//GatewayIntentGuildIntegrations contains events:
	//- GuildIntegrationsUpdate
	//- IntegrationCreate
	//- IntegrationUpdate
	//- IntegrationDelete
	GatewayIntentGuildIntegrations
	//GatewayIntentGuildWebhooks contains events:
	//- WebhooksUpdate
	GatewayIntentGuildWebhooks
	//GatewayIntentGuildInvites contains events:
	//- InviteCreate
	//- InviteDelete
	GatewayIntentGuildInvites
	//GatewayIntentGuildVoiceStates contains events:
	//- VoiceStateUpdate
	GatewayIntentGuildVoiceStates
	//GatewayIntentGuildPresences contains events:
	//- PresenceUpdate
	GatewayIntentGuildPresences
	//GatewayIntentGuildMessages contains events:
	//- MessageCreate
	//- MessageUpdate
	//- MessageDelete
	//- MessageDeleteBulk
	GatewayIntentGuildMessages
	//GatewayIntentGuildMessageReactions contains events:
	//- MessageReactionAdd
	//- MessageReactionRemove
	//- MessageReactionRemoveAll
	//- MessageReactionRemoveEmoji
	GatewayIntentGuildMessageReactions
	//GatewayIntentGuildMessageTyping contains events:
	//- TypingStart
	GatewayIntentGuildMessageTyping
	//GatewayIntentDirectMessages contains events:
	//- MessageCreate
	//- MessageUpdate
	//- MessageDelete
	//- ChannelPinsUpdate
	GatewayIntentDirectMessages
	//GatewayIntentDirectMessageReactions contains events:
	//- MessageReactionAdd
	//- MessageReactionRemove
	//- MessageReactionRemoveAll
	//- MessageReactionRemoveEmoji
	GatewayIntentDirectMessageReactions
	//GatewayIntentDirectMessageTyping contains events:
	//- TypingStart
	GatewayIntentDirectMessageTyping
	//GatewayIntentAll is a combination of all known GatewayIntents
	GatewayIntentAll GatewayIntent = (1 << (iota - 1)) - 1
)

//IsValid GatewayIntent
func (gatewayIntent GatewayIntent) IsValid() bool {
	return GatewayIntentAll&gatewayIntent == gatewayIntent && gatewayIntent != GatewayIntentNil
}

//Contains another GatewayIntent
func (gatewayIntent GatewayIntent) Contains(intent GatewayIntent) bool {
	return intent&gatewayIntent == intent && intent != GatewayIntentNil
}
