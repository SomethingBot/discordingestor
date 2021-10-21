package discordprimatives

import "time"

//SystemChannelFlag from https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
type SystemChannelFlag uint8

//ChannelType from https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType int

//VoiceQualityMode struct from https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
type VoiceQualityMode int

//Overwrite struct from https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
}

//User struct from https://discord.com/developers/docs/resources/user#user-object
type User struct {
}

//ThreadMetadata struct from https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
}

//ThreadMember struct from https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
}

//Channel struct from https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	//ID of Channel
	ID Snowflake `json:"id"`
	//Type of Channel
	Type ChannelType `json:"type"`
	//GuildID of Guild; may be missing for some channel objects received over gateway guild dispatches)
	GuildID Snowflake `json:"guild_id"`
	//Position of sorting
	Position int `json:"position"`
	//PermissionOverwrites for GuildMembers and Roles; these are explicit
	PermissionOverwrites []Overwrite `json:"permission_overwrites"`
	//Name of Channel (1-100 characters)
	Name string `json:"name"`
	//Topic of Channel (0-1024 characters)
	Topic string `json:"topic"`
	//IsNSFW Channel
	IsNSFW bool `json:"nsfw"`
	//LastMessageID for Channel (may not point to actual message)
	LastMessageID Snowflake `json:"last_message_id"`
	//Bitrate of Channel if VC
	Bitrate int `json:"bitrate"`
	//MemberLimit of Channel if VC
	MemberLimit int `json:"user_limit"`
	//RateLimitPerMember that a GuildMember has to wait before sending another message (0-21600); members with MANAGE_MESSAGES or MANAGE_CHANNEL are unaffected; applies to both message and create creation
	RateLimitPerMember int `json:"rate_limit_per_user"`
	//Recipients of a DM
	Recipients []User `json:"recipients"`
	//IconHash for grabbing Icon from CDN
	IconHash string `json:"icon"`
	//OwnerID of group DM or Thread
	OwnerID Snowflake `json:"owner_id"`
	//ApplicationID if group DM if bot-created
	ApplicationID Snowflake `json:"application_id"`
	//ParentID for a Guild Channel: category, Thread: Channel where created
	ParentID Snowflake `json:"parent_id"`
	//LastPinTimestamp maybe nil when no Messages have been pinned
	LastPinTimestamp time.Time `json:"last_pin_timestamp"`
	//VoiceRegion, empty if automatic
	VoiceRegion VoiceRegion `json:"rtc_region"`
	//VideoQualityMode of Channel, 1 when not present
	VoiceQualityMode VoiceQualityMode `json:"voice_quality_mode"`
	//MessageCount, this is Approximate, stops after 50
	MessageCount int `json:"message_count"`
	//MemberCount, this is Approximate, stops after 50
	MemberCount int `json:"member_count"`
	//ThreadMetadata not needed by other channels
	ThreadMetadata ThreadMetadata `json:"thread_metadata"`
	//Member aka current user if they have joined thread; only available for certain API Endpoints
	Member ThreadMember `json:"member"`
	//DefaultAutoArchiveDuration that clients use for newly made threads in minutes after activity; can be set to 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration"`
	//Permissions of Bot user including overwrites; only included when part of RESOLVED data received from SlashCommandInteraction
	Permissions string `json:"permissions"`
}
