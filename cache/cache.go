package cache

import (
	"errors"
	"github.com/RafikFarhad/hoax/config"
	"github.com/go-redis/redis/v8"
	"time"
)

type ProviderInterface interface {
	Get(key string, defaultValue string) (string, error)
	Push(key string, value string) error
	Delete(key string) error
	PushWithExpiry(key string, value string, duration time.Duration) error
	Ping() bool
}

type Cache struct {
	Name     string
	Provider ProviderInterface
}

var AppCache *Cache

func InitCache() error {
	cacheConfig := config.AppConfig.CacheConfig
	switch cacheConfig.Agent {
	case "redis":
		AppCache = &Cache{
			Name: "Redis",
			Provider: NewRedisCache(&redis.Options{
				Addr:     cacheConfig.RedisAddr,
				Password: cacheConfig.RedisPassword,
				DB:       cacheConfig.RedisDb,
			}),
		}
		break
	case "memory":
		AppCache = &Cache{
			Name:     "Memory",
			Provider: NewMemoryCache(),
		}
		break
	default:
		return errors.New("cache agent of type " + cacheConfig.Agent + " not yet supported")
	}
	if AppCache.Ping() == false {
		return errors.New("cache not responding")
	}
	return nil
}

func (c *Cache) Get(key string, defaultValue string) (string, error) {
	return c.Provider.Get(key, defaultValue)
}

func (c *Cache) Push(key string, value string) error {
	return c.Provider.Push(key, value)
}

func (c *Cache) Delete(key string) error {
	return c.Provider.Delete(key)
}

func (c *Cache) PushWithExpiry(key string, value string, duration time.Duration) error {
	return c.Provider.PushWithExpiry(key, value, duration)
}

func (c *Cache) Ping() bool {
	return c.Provider.Ping()
}
