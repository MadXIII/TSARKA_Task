package redis

import (
	"github.com/go-redis/redis"

	"github.com/madxiii/tsarka_task/configs"
)

func NewClients(cfg configs.Redis) (count *Count, hash *Hash) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	// to initialize key count in redis
	client.Set("count", 0, 0).Result()

	return NewCount(client), NewHash(client)
}
