package discordprimatives

import "time"

//GuildMember from https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
}

//WelcomeScreen from https://discord.com/developers/docs/resources/guild#welcome-screen-object
type WelcomeScreen struct {
}

//VerificationLevel from https://discord.com/developers/docs/resources/guild#guild-object-verification-level
type VerificationLevel uint8

//MessageNotificationsLevel from https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
type MessageNotificationsLevel uint8

//ExplicitContentFilterLevel from https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
type ExplicitContentFilterLevel uint8

//MFALevel from https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
type MFALevel uint8

//PremiumTier from https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
type PremiumTier uint8

//todo: fill documentation and IsValid
const (
	PremiumTier1 PremiumTier = iota + 1
	PremiumTier2
	PremiumTier3
)

//NSFWLevel from https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
type NSFWLevel uint8

//GuildFeature struct from https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature string

//Guild struct from https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	//ID of Guild
	ID Snowflake `json:"id,string"`
	//Name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
	//Icon hash for Icon
	IconHash ImageHash `json:"icon"`
	//IconHashInTemplate, returned when in the template object; todo: function for Fully qualified url
	IconHashInTemplate ImageHash `json:"icon_hash"`
	//SplashHash; todo: function for Fully qualified url
	SplashHash ImageHash `json:"splash"`
	//DiscoverySplashHash; only present for Guilds with the "DISCOVERABLE" feature; todo: function for Fully qualified url
	DiscoverySplashHash ImageHash `json:"discovery_splash"`
	//IsBotOwner is true if Bot is Guild owner, only sent when using the GET Current User Guilds endpoint and are relative to the requested user
	IsBotOwner bool `json:"owner"`
	//OwnerID of guild
	OwnerID Snowflake `json:"owner_id,string"`
	//Permissions of current user in Guild; total permissions for the Bot in the Guild (excludes overwrites)
	Permissions string `json:"permissions"`
	//RegionID; voice region ID for the Guild (deprecated)
	VoiceRegionID string `json:"region"`
	//AFKChannelID for Guild
	AFKChannelID Snowflake `json:"afk_channel_id,string"`
	//AFKTimeout in seconds
	AFKTimeout int `json:"afk_timeout"`
	//WidgetEnabled for Guild
	WidgetEnabled bool `json:"widget_enabled"`
	//WidgetChannelID that the widget will generate an Invite to, 0 if no invite
	WidgetChannelID Snowflake `json:"widget_channel_id,string"`
	//VerificationLevel required for the Guild
	VerificationLevel VerificationLevel `json:"verification_level"`
	//DefaultMessageNotificationsLevel for Guild
	DefaultMessageNotificationsLevel MessageNotificationsLevel `json:"default_message_notifications"`
	//ExplicitContentFilterLevel for Guild
	ExplicitContentFilterLevel ExplicitContentFilterLevel `json:"explicit_content_filter"`
	//Roles for Guild
	Roles []Role `json:"roles"`
	//Emojis is a list of custom Emojis
	Emojis []Emoji `json:"emojis"`
	//EnabledFeatures is a list of enabled GuildFeature(s)
	EnabledFeatures []GuildFeature `json:"features"`
	//MFALevel that is required for Guild
	MFALevel MFALevel `json:"mfa_level"`
	//ApplicationID of guild creator if bot-created
	ApplicationID Snowflake `json:"application_id,string"`
	//SystemChannelID is ID of Channel where Guild notices such as welcome message and boost events are posted
	SystemChannelID Snowflake `json:"system_channel_id,string"`
	//SystemChannelFlags for SystemChannel
	SystemChannelFlags SystemChannelFlag `json:"system_channel_flags,string"`
	//RulesChannelID where community Guilds can display rules and/or guidelines
	RulesChannelID Snowflake `json:"rules_channel_id,string"`
	//BotJoinedAt a timestamp when bot joined this Guild; only sent in GUILD_CREATE Event
	BotJoinedAt time.Time `json:"joined_at"`
	//IsLarge if Guild is large; only sent in GUILD_CREATE Event
	IsLarge bool `json:"large"`
	//IsUnavailable due to outage; only sent in GUILD_CREATE Event
	IsUnavailable bool `json:"unavailable"`
	//MemberCount in Guild; only sent in GUILD_CREATE Event
	MemberCount int `json:"member_count"`
	//VoiceStates is a list of VoiceState; only sent in GUILD_CREATE Event; VoiceState(s) lack Guild ID
	VoiceStates []VoiceState `json:"voice_states"`
	//Members in Guild; only sent in GUILD_CREATE Event
	Members []GuildMember `json:"members"`
	//Channels in Guild; only sent in GUILD_CREATE Event
	Channels []Channel `json:"channels"`
	//Threads that Bot has permission to view; only sent in GUILD_CREATE Event
	Threads []Channel `json:"threads"`
	//Presences in Guild; only sent in GUILD_CREATE Event; only includes non-offline GuildMember(s) if IsLarge
	Presences []PresenceUpdate `json:"presences"`
	//MaxPresences in Guild; almost always 0 unless Guild is massive
	MaxPresences int `json:"max_presences"`
	//MaxMembers in Guild
	MaxMembers int `json:"max_members"`
	//VanityUrlCode for Guild
	VanityUrlCode string `json:"vanity_url_code"`
	//Description of a Community Guild
	Description string `json:"description"`
	//BannerHash for Guild
	BannerHash ImageHash `json:"banner"`
	//PremiumTier aka boost level
	PremiumTier PremiumTier `json:"premium_tier"`
	//PremiumSubscriptionCount is number of boosts Guild has
	PremiumSubscriptionCount int `json:"premium_subscription_count"`
	//PreferredLocale of Guild; used for Guild Discovery and Discord notices; defaults to en-US
	PreferredLocale string `json:"preferred_locale"`
	//PublicUpdatesChannelID where Guilds get Discord notices
	PublicUpdatesChannelID Snowflake `json:"public_updates_channel_id"`
	//MaxVideoChannelUsers in a Channel
	MaxVideoChannelUsers int `json:"max_video_channel_users"`
	//ApproximateMemberCount; returned from GET /guilds/<id> endpoint when with_counts is true
	ApproximateMemberCount int `json:"approximate_member_count"`
	//WelcomeScreen of a Community Guild, shown to new members, returned in an Invite
	WelcomeScreen WelcomeScreen `json:"welcome_screen"`
	//NSFWLevel of Guild
	NSFWLevel NSFWLevel `json:"nsfw_level"`
	//StageInstances in Guild; only sent in GUILD_CREATE Event
	StageInstances []StageInstance `json:"stage_instances"`
	//Stickers in Guild
	Stickers []Sticker `json:"stickers"`
}
