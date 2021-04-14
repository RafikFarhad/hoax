package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	ProviderInterface
	redisClient *redis.Client
}

func NewRedisCache(options *redis.Options) RedisCache {
	return RedisCache{
		redisClient: redis.NewClient(options),
	}
}

func (r RedisCache) Get(key string, defaultValue string) (string, error) {
	val, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		val = defaultValue
	}
	return val, err
}

func (r RedisCache) Push(key string, value string) error {
	return r.redisClient.Set(context.Background(), key, value, 0).Err()
}

func (r RedisCache) Delete(key string) error {
	return r.redisClient.Del(context.Background(), key).Err()
}

func (r RedisCache) PushWithExpiry(key string, value string, duration time.Duration) error {
	return r.redisClient.SetEX(context.Background(), key, value, duration).Err()
}

func (r RedisCache) Ping() bool {
	err := r.PushWithExpiry("__test__", "ok", time.Minute)
	return err == nil
}
