package db

import "github.com/go-redis/redis"
import "fmt"
import "errors"

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

//QueryWithRedis ...
func QueryWithRedis(s func(*redis.Client) error) error {
	if client != nil {
		return s(client)
	}
	return errors.New("redis not connected")
}
