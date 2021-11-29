package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6397",
		Password: "",
		DB:       0,
	})
	store := &RedisStore{
		client: client,
	}
	return store
}

func (r *RedisStore) Set(key, value string, second int) {
	r.client.Set(key, value, time.Second*time.Duration(second))
}

func (r *RedisStore) Get(key string) (value string, err error) {
	return r.client.Get(key).Result()
}
