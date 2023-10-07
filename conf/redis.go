package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var redisClient *redis.Client

var DefaultDuration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, err
}

func (rc *RedisClient) Set(key string, value any, rest ...any) error {
	nd := DefaultDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			nd = v
		}
	}
	return redisClient.Set(context.Background(), key, value, nd).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return redisClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return redisClient.Del(context.Background(), key...).Err()
}

func (rc *RedisClient) GetExpireDuration(key string) (time.Duration, error) {
	return redisClient.TTL(context.Background(), key).Result()
}
