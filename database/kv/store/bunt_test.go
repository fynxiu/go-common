package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuntStore_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		preFn   func(store Store)
	}{
		{
			name: "key exists",
			args: args{
				key: "key1",
			},
			want:    "value1",
			wantErr: false,
			preFn: func(store Store) {
				store.Set("key1", "value1", time.Hour)
			},
		},
		{
			name: "key not exists",
			args: args{
				key: "key1",
			},
			wantErr: true,
		},
		{
			name: "key expired",
			args: args {
				key: "key",
			},
			wantErr: true,
			preFn: func(store Store) {
				store.Set("key", "value", time.Nanosecond)
				time.Sleep(time.Nanosecond)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewBuntStore()
			assert.NoError(t, err)
			if tt.preFn != nil {
				tt.preFn(s)
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuntStore_Set(t *testing.T) {

	var key = "key"
	store, err := NewBuntStore()
	assert.NoError(t, err)
	store.Set(key, "value", time.Hour)
	v , _ := store.Get(key)
	assert.Equal(t, "value", v)

	store.Set(key, "value_new", time.Hour)
	v , _ = store.Get(key)
	assert.Equal(t, "value_new", v)
}
