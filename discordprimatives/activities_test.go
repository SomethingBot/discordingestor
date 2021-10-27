package discordprimatives

import "testing"

func TestPresenceStatus_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		status PresenceStatus
		want   bool
	}{
		{
			name:   "PresenceStatusIdle",
			status: "idle",
			want:   true,
		},
		{
			name:   "PresenceStatusDnd",
			status: "dnd",
			want:   true,
		},
		{
			name:   "PresenceStatusOnline",
			status: "offline",
			want:   true,
		},
		{
			name:   "PresenceStatusNil",
			status: "",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityType_IsValid(t *testing.T) {
	tests := []struct {
		name         string
		activityType ActivityType
		want         bool
	}{
		{
			name:         "ActivityTypeGame",
			activityType: ActivityTypeGame,
			want:         true,
		},
		{
			name:         "ActivityTypeStreaming",
			activityType: ActivityTypeStreaming,
			want:         true,
		},
		{
			name:         "ActivityTypeListening",
			activityType: ActivityTypeListening,
			want:         true,
		},
		{
			name:         "ActivityTypeWatching",
			activityType: ActivityTypeWatching,
			want:         true,
		},
		{
			name:         "ActivityTypeCustom",
			activityType: ActivityTypeCustom,
			want:         true,
		},
		{
			name:         "ActivityTypeCompeting",
			activityType: ActivityTypeCompeting,
			want:         true,
		},
		{
			name:         "ActivityTypeINVALID",
			activityType: 255,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.activityType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityFlags_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		activityFlags ActivityFlags
		want          bool
	}{
		{
			name:          "ActivityFlagInstance",
			activityFlags: ActivityFlagInstance,
			want:          true,
		},
		{
			name:          "ActivityFlagJoin",
			activityFlags: ActivityFlagJoin,
			want:          true,
		},
		{
			name:          "ActivityFlagSpectate",
			activityFlags: ActivityFlagSpectate,
			want:          true,
		},
		{
			name:          "ActivityFlagJoinRequest",
			activityFlags: ActivityFlagJoinRequest,
			want:          true,
		},
		{
			name:          "ActivityFlagSync",
			activityFlags: ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagPlay",
			activityFlags: ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagALL",
			activityFlags: ActivityFlagALL,
			want:          false,
		},
		{
			name:          "ActivityFlagSync|ActivityFlagPlay",
			activityFlags: ActivityFlagSync | ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagNil",
			activityFlags: ActivityFlagNil,
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.activityFlags.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityFlags_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		activityFlags ActivityFlags
		want          bool
	}{
		{
			name:          "ActivityFlagInstance",
			activityFlags: ActivityFlagInstance,
			want:          true,
		},
		{
			name:          "ActivityFlagJoin",
			activityFlags: ActivityFlagJoin,
			want:          true,
		},
		{
			name:          "ActivityFlagSpectate",
			activityFlags: ActivityFlagSpectate,
			want:          true,
		},
		{
			name:          "ActivityFlagJoinRequest",
			activityFlags: ActivityFlagJoinRequest,
			want:          true,
		},
		{
			name:          "ActivityFlagSync",
			activityFlags: ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagPlay",
			activityFlags: ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagNil",
			activityFlags: ActivityFlagNil,
			want:          false,
		},
		{
			name:          "ActivityFlagPlay|ActivityFlagSync",
			activityFlags: ActivityFlagPlay | ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagALL",
			activityFlags: ActivityFlagALL,
			want:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ActivityFlagALL.Contains(tt.activityFlags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}