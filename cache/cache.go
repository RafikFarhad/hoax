package cache

import (
	"errors"
	"github.com/RafikFarhad/hoax/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type BaseCache interface {
	Get(key string, defaultValue string) (string, error)
	Push(key string, value string) error
	Delete(key string) error
	PushWithExpiry(key string, value string, duration time.Duration) error
	Ping() bool
}

var Cache BaseCache

func InitCache() error {
	cacheConfig := config.AppConfig.CacheConfig
	switch cacheConfig.Agent {
	case "":
		return nil
	case "redis":
		Cache = NewRedisCache(&redis.Options{
			Addr:     cacheConfig.RedisAddr,
			Password: cacheConfig.RedisPassword,
			DB:       cacheConfig.RedisDb,
		})
		break
	default:
		return errors.New("cache agent of type " + cacheConfig.Agent + " not yet supported")
	}
	if Cache.Ping() == false {
		return errors.New("cache not responding")
	}
	return nil
}
