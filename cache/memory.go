package cache

import (
	"errors"
	"time"
)

type MemoryCache struct {
	ProviderInterface
	store map[string]struct {
		string
		int64
	}
}

func NewMemoryCache() MemoryCache {
	return MemoryCache{
		store: make(map[string]struct {
			string
			int64
		}),
	}
}

func (mc MemoryCache) Get(key string, defaultValue string) (string, error) {
	val, ok := mc.store[key]
	if ok == false {
		return defaultValue, errors.New("cache missed")
	}
	if val.int64 == 0 || val.int64 > time.Now().Unix() {
		return val.string, nil
	}
	return defaultValue, nil
}

func (mc MemoryCache) Push(key string, value string) error {
	mc.store[key] = struct {
		string
		int64
	}{string: value, int64: 0}
	return nil
}

func (mc MemoryCache) Delete(key string) error {
	delete(mc.store, key)
	return nil
}

func (mc MemoryCache) PushWithExpiry(key string, value string, duration time.Duration) error {
	mc.store[key] = struct {
		string
		int64
	}{string: value, int64: time.Now().Add(duration).Unix()}
	return nil
}

func (mc MemoryCache) Ping() bool {
	err1 := mc.Push("__test__", "__success__")
	val, err2 := mc.Get("__test__", "")
	return err1 == nil && val == "__success__" && err2 == nil
}
