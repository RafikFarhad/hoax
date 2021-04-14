package config

type CacheConfig struct {
	Agent string `ini:"cache_agent"`

	// redis cache
	RedisAddr     string `ini:"redis_addr"`
	RedisPassword string `ini:"redis_password"`
	RedisDb       int    `ini:"redis_db"`
}

func NewCacheConfig() *CacheConfig {
	return &CacheConfig{
		Agent: "memory",
	}
}
