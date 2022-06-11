package redis

import (
	"github.com/go-redis/redis"
)

type Hash struct {
	client *redis.Client
}

func NewHash(client *redis.Client) *Hash {
	return &Hash{client: client}
}

func (r *Hash) StoreKey(key string) {
	r.client.Set(key, 0, 0)
}

func (r *Hash) StoreValByKey(key string, val int) {
	r.client.Set(key, val, 0)
}

func (r *Hash) GetValueByKey(key string) (int, error) {
	res, err := r.client.Get(key).Int()
	return res, err
}
