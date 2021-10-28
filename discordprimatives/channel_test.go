package discordprimatives

import "testing"

func TestSystemChannelFlag_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		systemChannelFlag SystemChannelFlag
		want              bool
	}{
		{
			name:              "SystemChannelFlagNil",
			systemChannelFlag: SystemChannelFlagNil,
			want:              false,
		},
		{
			name:              "SystemChannelFlagSuppressJoinNotifications",
			systemChannelFlag: SystemChannelFlagSuppressJoinNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications",
			systemChannelFlag: SystemChannelFlagSuppressGuildReminderNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagALL",
			systemChannelFlag: SystemChannelFlagALL,
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.systemChannelFlag.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemChannelFlag_Contains(t *testing.T) {
	tests := []struct {
		name              string
		systemChannelFlag SystemChannelFlag
		flags             SystemChannelFlag
		want              bool
	}{
		{
			name:              "SystemChannelFlagNil",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagNil,
			want:              false,
		},
		{
			name:              "SystemChannelFlagSuppressJoinNotifications",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressJoinNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressGuildReminderNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagALL",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagALL,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications|SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressGuildReminderNotifications | SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.systemChannelFlag.Contains(tt.flags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
