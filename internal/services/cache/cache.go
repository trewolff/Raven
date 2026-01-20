package cache

import (
	"context"
	"raven/internal/config"

	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	GetResultByKey(key string) (string, error)
}

type DefaultCacheService struct {
	cache *redis.Client
}

func NewCacheConnection() *redis.Client {
	conf, _ := config.GetConfig()
	return redis.NewClient(&redis.Options{
		Addr:     conf.REDIS_URL,
		Password: conf.REDIS_PASSWORD,
	})
}

func NewCacheService(cache *redis.Client) CacheService {
	return &DefaultCacheService{
		cache: cache,
	}
}

var ctx = context.Background()

func (c *DefaultCacheService) GetResultByKey(key string) (string, error) {
	val, err := c.cache.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
