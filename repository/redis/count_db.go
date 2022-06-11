package redis

import (
	"github.com/go-redis/redis"
)

type Count struct {
	client *redis.Client
}

func NewCount(client *redis.Client) *Count {
	return &Count{client: client}
}

func (c *Count) Value() (string, error) {
	res, err := c.client.Get("count").Result()
	return res, err
}

func (c *Count) Add(num int) error {
	res, err := c.client.Get("count").Int()
	if err != nil {
		return err
	}

	res += num

	_, err = c.client.Set("count", res, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *Count) Sub(num int) error {
	res, err := c.client.Get("count").Int()
	if err != nil {
		return err
	}

	res -= num

	_, err = c.client.Set("count", res, 0).Result()
	if err != nil {
		return err
	}

	return nil
}
