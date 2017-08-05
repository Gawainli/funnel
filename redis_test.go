package funnel

import (
	"testing"

	"github.com/go-redis/redis"
)

func Test_RedisConnect(t *testing.T) {
	err := ConnectRedis("192.168.0.98:6379", "")
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_RedisSet(t *testing.T) {
	Test_RedisConnect(t)
	query := func(c *redis.Client) error {
		return c.Set("test:test:1", "test1", 0).Err()
	}
	QueryWithRedis(query)
}
