package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

//ConnectRedis ...
func ConnectRedis(addr string, pwd string) error {
	options := &redis.Options{}
	options.Addr = addr
	options.Password = pwd
	client = redis.NewClient(options)
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return err
}
