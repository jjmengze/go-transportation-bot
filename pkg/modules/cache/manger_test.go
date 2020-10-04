package cache

import (
	"reflect"
	"testing"
)

func TestGetManager(t *testing.T) {
	tests := []struct {
		name string
		want *Manager
	}{
		{
			name: "fill Manger instance",
			want: &Manager{
				RedisConnections: make(map[string]*redisClientHolder),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManager() = %v, want %v", got, tt.want)
			}
		})
	}
}