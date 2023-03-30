package config

import (
	"time"
)

var store Store

func init() {
	if store == nil {
		store = NewSyncMapStore()
	}
}

type Store interface {
	GetInt(key string) int
	GetBool(key string) bool
	GetString(key string) string
	GetStringMap(key string) map[string]any
	GetDuration(key string) time.Duration
	Set(key string, val any)
	SetDefault(key string, val any)
	BindEnv(input ...string) error
	GetConfig() any
}

var hooks []func()

func AddStoreSetHook(hook func()) {
	hooks = append(hooks, hook)
}

func ApplyStoreSetHooks() {
	for _, hook := range hooks {
		hook()
	}
}

func GetStore() Store {
	return store
}

func SetStore(newStore Store) {
	store = newStore
	ApplyStoreSetHooks()
}

func GetInt(key string) int {
	return store.GetInt(key)
}

func GetBool(key string) bool {
	return store.GetBool(key)
}

func GetString(key string) string {
	return store.GetString(key)
}

func GetStringMap(key string) map[string]any {
	return store.GetStringMap(key)
}

func GetDuration(key string) time.Duration {
	return store.GetDuration(key)
}
func Set(key string, val any) {
	store.Set(key, val)
}

func SetDefault(key string, val any) {
	store.SetDefault(key, val)
}

func GetConfig() any {
	return store.GetConfig()
}

func Debug(mod ...string) bool {
	if len(mod) == 0 {
		return store.GetBool("debug")
	}
	return store.GetBool("debug") || store.GetBool(mod[0]+".debug")
}

func BindEnv(input ...string) error {
	return store.BindEnv(input...)
}
