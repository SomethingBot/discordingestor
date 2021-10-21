package ingestor

//todo: check if types need a specific uint (uint32/64)
//todo: change all references from a server to a Guild

type DiscordIntent uint32
type Snowflake uint32

//Emoji struct from json, documented at https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
}

//Role struct from json, documented at https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
}

//GuildFeature struct from json, documented at https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature struct {
}

//VoiceState struct from json, documented at https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
}

//Channel struct from json, documented at https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
}

//GuildMember struct from json, documented at https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
}

//Presence struct from json, documented at https://discord.com/developers/docs/topics/gateway#presence-update
type Presence struct {
}

//WelcomeScreen struct from json, documented at https://discord.com/developers/docs/resources/guild#welcome-screen-object
type WelcomeScreen struct {
}

//StageInstance struct from json, documented at https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
}

//Sticker struct from json, documented at https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
}

//Guild struct from json, documented at https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	//ID of Guild
	ID string `json:"id"`
	//Name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
	//Icon hash; todo: function for Fully qualified url
	IconHash string `json:"icon"`
	//IconHashInTemplate, returned when in the template object; todo: function for Fully qualified url
	IconHashInTemplate string `json:"icon_hash?"`
	//SplashHash; todo: function for Fully qualified url
	SplashHash string `json:"splash"`
	//DiscoverySplash hash; only present for Guilds with the "DISCOVERABLE" feature; todo: function for Fully qualified url
	DiscoverySplash string `json:"discovery_splash"`
	//IsBotOwner is true if Bot is Guild owner, only sent when using the GET Current User Guilds endpoint and are relative to the requested user
	IsBotOwner bool `json:"owner?"`
	//OwnerID of guild
	OwnerID Snowflake `json:"owner_id"`
	//Permissions of current user in Guild; total permissions for the Bot in the Guild (excludes overwrites)
	Permissions string `json:"permissions?"`
	//RegionID; voice region ID for the Guild (deprecated)
	RegionID string `json:"region?"`
	//AFKChannelID for Guild
	AFKChannelID Snowflake `json:"afk_channel_id"`
	//AFKTimeout in seconds
	AFKTimeout int `json:"afk_timeout"`
	//WidgetEnabled for Guild
	WidgetEnabled bool `json:"widget_enabled?"`
	//WidgetChannelID that the widget will generate an Invite to, 0 if no invite
	WidgetChannelID Snowflake `json:"widget_channel_id?"`
	//VerificationLevel required for the Guild
	VerificationLevel int `json:"verification_level"`
	//DefaultMessageNotificationsLevel for Guild
	DefaultMessageNotificationsLevel int `json:"default_message_notifications"`
	//ExplicitContentFilterLevel for Guild
	ExplicitContentFilterLevel int `json:"explicit_content_filter"`
	//Roles for Guild
	Roles []Role `json:"roles"`
	//Emojis is a list of custom Emojis
	Emojis []Emoji `json:"emojis"`
	//EnabledFeatures is a list of enabled GuildFeature(s)
	EnabledFeatures []GuildFeature `json:"features"`
	//MFALevel that is required for Guild
	MFALevel int `json:"mfa_level"`
	//ApplicationID of guild creator if bot-created
	ApplicationID Snowflake `json:"application_id"`
	//SystemChannelID is ID of Channel where Guild notices such as welcome message and boost events are posted
	SystemChannelID Snowflake `json:"system_channel_id"`
	//SystemChannelFlags for SystemChannel
	SystemChannelFlags int `json:"system_channel_flags"`
	//RulesChannelID where community Guilds can display rules and/or guidelines
	RulesChannelID Snowflake `json:"rules_channel_id"`
	//todo: maybe convert this to a type that doesn't have to be parsed first
	//BotJoinedAt a ISO8601 timestamp when bot joined this Guild; only sent in GUILD_CREATE Event
	BotJoinedAt string `json:"joined_at?"`
	//IsLarge if Guild is large; only sent in GUILD_CREATE Event
	IsLarge bool `json:"large?"`
	//IsUnavailable due to outage; only sent in GUILD_CREATE Event
	IsUnavailable bool `json:"unavailable?"`
	//MemberCount in Guild; only sent in GUILD_CREATE Event
	MemberCount int `json:"member_count?"`
	//VoiceStates is a list of VoiceState; only sent in GUILD_CREATE Event; VoiceState(s) lack Guild ID
	VoiceStates []VoiceState `json:"voice_states?"`
	//Members in Guild; only sent in GUILD_CREATE Event
	Members []GuildMember `json:"members?"`
	//Channels in Guild; only sent in GUILD_CREATE Event
	Channels []Channel `json:"channels?"`
	//Threads that Bot has permission to view; only sent in GUILD_CREATE Event
	Threads []Channel `json:"threads?"`
	//Presences in Guild; only sent in GUILD_CREATE Event; only includes non-offline GuildMember(s) if IsLarge
	Presences []Presence `json:"presences?"`
	//MaxPresences in Guild; almost always 0 unless Guild is massive
	MaxPresences int `json:"max_presences?"`
	//MaxMembers in Guild
	MaxMembers int `json:"max_members?"`
	//VanityUrlCode for Guild
	VanityUrlCode string `json:"vanity_url_code"`
	//Description of a Community Guild
	Description string `json:"description?"`
	//BannerHash todo: function for Fully qualified url
	BannerHash string `json:"banner?"`
	//PremiumTier aka boost level
	PremiumTier int `json:"premium_tier"`
	//PremiumSubscriptionCount is number of boosts Guild has
	PremiumSubscriptionCount int `json:"premium_subscription_count?"`
	//PreferredLocale of Guild; used for Guild Discovery and Discord notices; defaults to en-US
	PreferredLocale string `json:"preferred_locale"`
	//PublicUpdatesChannelID where Guilds get Discord notices
	PublicUpdatesChannelID Snowflake `json:"public_updates_channel_id"`
	//MaxVideoChannelUsers in a Channel
	MaxVideoChannelUsers int `json:"max_video_channel_users?"`
	//ApproximateMemberCount; returned from GET /guilds/<id> endpoint when with_counts is true
	ApproximateMemberCount int `json:"approximate_member_count?"`
	//WelcomeScreen of a Community Guild, shown to new members, returned in an Invite
	WelcomeScreen WelcomeScreen `json:"welcome_screen?"`
	//NSFWLevel of Guild
	NSFWLevel int `json:"nsfw_level"`
	//StageInstances in Guild; only sent in GUILD_CREATE Event
	StageInstances []StageInstance `json:"stage_instances?"`
	//Stickers in Guild
	Stickers []Sticker `json:"stickers?"`
}