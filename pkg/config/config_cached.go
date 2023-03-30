package config

import (
	"sync"
	"time"
)

func Cached(store Store) Store {
	layer := store.GetConfig()
	if _, ok := layer.(*sync.Map); ok {
		return store
	} else {
		return &plainCachedStore{Map: &sync.Map{}, store: store}
	}
}

type plainCachedStore struct {
	*sync.Map
	store Store
}

func (store *plainCachedStore) GetInt(key string) int {
	if val, ok := store.Load(key); ok {
		return val.(int)
	} else {
		val := store.store.GetInt(key)
		store.Store(key, val)
		return val
	}
}

func (store *plainCachedStore) GetBool(key string) bool {
	if val, ok := store.Load(key); ok {
		return val.(bool)
	} else {
		val := store.store.GetBool(key)
		store.Store(key, val)
		return val
	}
}

func (store *plainCachedStore) GetString(key string) string {
	if val, ok := store.Load(key); ok {
		return val.(string)
	} else {
		val := store.store.GetString(key)
		store.Store(key, val)
		return val
	}
}

func (store *plainCachedStore) GetStringMap(key string) map[string]any {
	if val, ok := store.Load(key + "."); ok {
		return val.(map[string]any)
	} else {
		val := store.store.GetStringMap(key)
		store.Store(key+".", val)
		return val
	}
}

func (store *plainCachedStore) GetDuration(key string) time.Duration {
	if val, ok := store.Load(key); ok {
		return val.(time.Duration)
	} else {
		val := store.store.GetDuration(key)
		store.Store(key, val)
		return val
	}
}

func (store *plainCachedStore) Set(key string, val any) {
	store.Map.Delete(key)
	store.store.Set(key, val)
}

func (store *plainCachedStore) SetDefault(key string, val any) {
	store.Map.Delete(key)
	store.store.SetDefault(key, val)
}

func (store *plainCachedStore) GetConfig() any {
	return store.Map
}
func (store *plainCachedStore) BindEnv(input ...string) error {
	return store.store.BindEnv(input...)
}
