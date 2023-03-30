package config

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func NewSyncMapStore() Store {
	return &syncMapStore{&sync.Map{}}
}

type syncMapStore struct {
	*sync.Map
}

func (store *syncMapStore) GetInt(key string) int {
	if val, ok := store.Load(key); ok {
		return val.(int)
	} else {
		return 0
	}
}

func (store *syncMapStore) GetBool(key string) bool {
	if val, ok := store.Load(key); ok {
		return val.(bool)
	} else {
		return false
	}
}

func (store *syncMapStore) GetString(key string) string {
	if val, ok := store.Load(key); ok {
		return val.(string)
	} else {
		return ""
	}
}

func (store *syncMapStore) GetStringMap(key string) map[string]any {
	var out = make(map[string]any)
	store.Map.Range(func(k, v interface{}) bool {
		if str, ok := k.(string); ok {
			if strings.HasPrefix(str, key+".") {
				out[strings.TrimPrefix(str, key+".")] = v
			}
		}
		return true
	})
	return out
}

func (store *syncMapStore) GetDuration(key string) time.Duration {
	if val, ok := store.Load(key); ok {
		return val.(time.Duration)
	} else {
		return 0
	}
}

func (store *syncMapStore) Set(key string, val any) {
	store.Store(key, val)
}

func (store *syncMapStore) SetDefault(key string, val any) {
	store.Store(key, val)
}

func (store *syncMapStore) GetConfig() any {
	return store.Map
}

func (store *syncMapStore) BindEnv(input ...string) error {
	if len(input) == 0 {
		return fmt.Errorf("missing key to bind to")
	}

	key := strings.ToLower(input[0])

	if len(input) == 1 {
		store.Store(key, os.Getenv(input[0]))
	} else {
		store.Store(input[0], os.Getenv(input[1]))
	}

	return nil
}
