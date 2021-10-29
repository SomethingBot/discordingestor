package discordprimatives

import (
	"time"
)

//Snowflake documented at https://discord.com/developers/docs/reference#snowflakes
type Snowflake uint64 //todo: helper methods

const (
	DiscordEpoch int64 = 1420070400000
)

//IsValid Snowflake
func (snowflake Snowflake) IsValid() bool {
	return snowflake != 0
}

//Timestamp of Snowflake as a time.Time UTC
func (snowflake Snowflake) Timestamp() time.Time {
	snowflakeTimeMilli := int64(snowflake>>22) + DiscordEpoch
	return time.Unix(snowflakeTimeMilli/1e3, (snowflakeTimeMilli%1e3)*1e6).UTC()
}

//InternalWorkerID used by Discord
func (snowflake Snowflake) InternalWorkerID() uint8 {
	return uint8((snowflake & 0x3E0000) >> 17)
}

//InternalProcessID used by Discord
func (snowflake Snowflake) InternalProcessID() uint8 {
	return uint8((snowflake & 0x1F000) >> 12)
}

//Increment (ed) for every ID generated on InternalProcessID
func (snowflake Snowflake) Increment() uint16 {
	return uint16(snowflake & 0xFFF)
}
