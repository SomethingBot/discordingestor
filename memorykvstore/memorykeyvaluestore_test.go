package memorykvstore

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want *MemoryKeyValueStore
	}{
		{
			name: "ValidStore",
			want: &MemoryKeyValueStore{
				store: make(map[string][]byte),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryKeyValueStore_Open(t *testing.T) {
	t.Parallel()
	if err := New().Open(); err != nil {
		t.Fatalf("New() returned non-nil error (%v)\n", err)
	}
}

func TestMemoryKeyValueStore_Close(t *testing.T) {
	t.Parallel()
	if err := New().Close(); err != nil {
		t.Fatalf("New() returned non-nil error (%v)\n", err)
	}
}

func TestMemoryKeyValueStore_GetSet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		key     string
		set     string
		get     string
		wantErr bool
	}{
		{
			name:    "opus",
			key:     "opus",
			set:     "opus",
			get:     "opus",
			wantErr: false,
		},
		{
			name:    "boom",
			key:     "boomkey",
			set:     "boomset",
			get:     "boom1",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Parallel()
			kvs := New()
			err := kvs.Set(test.key, strings.NewReader(test.set))
			if err != nil {
				if !test.wantErr {
					t.Fatalf("wanted no error on set, got error (%v)\n", err)
				}
			}
			get, err := kvs.Get(test.key)
			if err != nil {
				if !test.wantErr {
					t.Fatalf("wanted no error on get, got error (%v)\n", err)
				}
			}
			getBytes, err := io.ReadAll(get)
			if err != nil {
				if !test.wantErr {
					t.Fatalf("wanted no error on readall, got error (%v)\n", err)
				}
			}
			if string(getBytes) != test.get && !test.wantErr {
				t.Fatalf("set (%v), but got back (%v)\n", test.set, string(getBytes))
			}
		})
	}
}
