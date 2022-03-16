package primitives

import (
	"math"
	"time"
)

//MessageType documented at https://discord.com/developers/docs/resources/channel#message-object-message-types
type MessageType uint8

const (
	//MessageTypeDefault is the Default MessageType
	MessageTypeDefault MessageType = iota
	//MessageTypeRecipientAdd is the MessageType when a User is added to a Channel
	MessageTypeRecipientAdd
	//MessageTypeRecipientRemove is the MessageType when a User is removed from a Channel
	MessageTypeRecipientRemove
	//MessageTypeCall is when a MessageType is a Call
	MessageTypeCall
	//MessageTypeChannelNameChange is when a Channel Name is changed
	MessageTypeChannelNameChange
	//MessageTypeChannelIconChange is when a Channel Icon is changed
	MessageTypeChannelIconChange
	//MessageTypeChannelPinnedMessage is when a Channel message is changed
	MessageTypeChannelPinnedMessage
	//MessageTypeGuildMemberJoin is when a User is added to a channel
	MessageTypeGuildMemberJoin
	//MessageTypeUserPremiumGuildSubscription is when a GuildMember boosts a Guild
	MessageTypeUserPremiumGuildSubscription
	//MessageTypeUserPremiumGuildSubscriptionTier1 is when a Guild reaches PremiumTier1
	MessageTypeUserPremiumGuildSubscriptionTier1
	//MessageTypeUserPremiumGuildSubscriptionTier2 is when a Guild reaches PremiumTier2
	MessageTypeUserPremiumGuildSubscriptionTier2
	//MessageTypeUserPremiumGuildSubscriptionTier3 is when a Guild reaches PremiumTier3
	MessageTypeUserPremiumGuildSubscriptionTier3
	//MessageTypeChannelFollowAdd is when a User adds another ChannelTypeGuildNews to a Channel
	MessageTypeChannelFollowAdd
	//MessageTypeGuildDiscoveryGracePeriodInitialWarning is when a Guild is about to lose its spot in Guild Discovery
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	//MessageTypeGuildDiscoveryGracePeriodFinalWarning is when a Guild is imminent about to lose its spot in Guild Discovery
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	//MessageTypeThreadCreated is when a Thread has been created
	MessageTypeThreadCreated
	//MessageTypeReply is when a Message is a reply to another Message
	MessageTypeReply
	//MessageTypeChatInputCommand is when a message is a slash command? maybe? todo: what is this message type?
	MessageTypeChatInputCommand
	//MessageTypeThreadStarterMessage is the message that starts a Thread
	MessageTypeThreadStarterMessage
	//MessageTypeGuildInviteReminder is the message that a Guild should add an Invite
	MessageTypeGuildInviteReminder
	//MessageTypeContextMenuCommand is part of a slash command?
	MessageTypeContextMenuCommand
	//MessageTypeInvalid is an intentionally invalid MessageType
	MessageTypeInvalid MessageType = math.MaxUint8
)

//IsValid MessageType
func (messageType MessageType) IsValid() bool {
	switch messageType {
	case MessageTypeDefault,
		MessageTypeRecipientAdd,
		MessageTypeRecipientRemove,
		MessageTypeCall,
		MessageTypeChannelNameChange,
		MessageTypeChannelIconChange,
		MessageTypeChannelPinnedMessage,
		MessageTypeGuildMemberJoin,
		MessageTypeUserPremiumGuildSubscription,
		MessageTypeUserPremiumGuildSubscriptionTier1,
		MessageTypeUserPremiumGuildSubscriptionTier2,
		MessageTypeUserPremiumGuildSubscriptionTier3,
		MessageTypeChannelFollowAdd,
		MessageTypeGuildDiscoveryGracePeriodInitialWarning,
		MessageTypeGuildDiscoveryGracePeriodFinalWarning,
		MessageTypeThreadCreated,
		MessageTypeReply,
		MessageTypeChatInputCommand,
		MessageTypeThreadStarterMessage,
		MessageTypeGuildInviteReminder,
		MessageTypeContextMenuCommand:
		return true
	default:
		return false
	}
}

//Attachment documented at https://discord.com/developers/docs/resources/channel#attachment-object
type Attachment struct {
	//ID of the Attachment
	ID Snowflake `json:"id"`
	//FileName of the Attachment
	FileName string `json:"filename"`
	//Description for the File
	Description string `json:"description"`
	//ContentType is the MIME type of the Attachment
	ContentType string `json:"content_type"`
	//Size of the File in bytes
	Size int `json:"size"`
	//URL for File
	URL string `json:"URL"`
	//ProxyURL of a proxied Attachment
	ProxyURL string `json:"proxy_url"`
	//Height of File (if image)
	Height int `json:"height"`
	//Width of File (if image)k
	Width int `json:"width"`
	//IsEphemeral Attachment that will be removed after set period of time
	IsEphemeral bool `json:"ephemeral"`
}

//Message documented at https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                Snowflake        `json:"id"`
	ChannelID         Snowflake        `json:"channel_id"`
	GuildID           Snowflake        `json:"guild_id"`
	Author            User             `json:"author"`
	Member            GuildMember      `json:"member"`
	Content           string           `json:"content"`
	Timestamp         time.Time        `json:"timestamp"`
	EditedTimestamp   time.Time        `json:"edited_timestamp"`
	IsTextToSpeech    bool             `json:"tts"`
	IsMentionEveryone bool             `json:"mention_everyone"`
	Mentions          []User           `json:"mentions"`
	MentionRoles      []Role           `json:"mention_roles"`
	MentionChannels   []ChannelMention `json:"mention_channels"`
	Attachments       []Attachment     `json:"attachments"`
	Embeds            []Embed          `json:"embeds"`
}
