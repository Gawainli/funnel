package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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

//QueryWithRedis ...
func QueryWithRedis(s func(*redis.Client) error) error {
	if client != nil {
		return s(client)
	}
	return errors.New("redis not connected")
}

//KeyLock ...
func KeyLock(key string) {
	keylock := key + ":lock"
	redisQuery := func(r *redis.Client) error {
		err := r.SetNX(keylock, strconv.FormatInt(time.Now().Unix(), 10), 3*time.Second).Err()
		return err
	}
	err := QueryWithRedis(redisQuery)
	for err != nil {
		time.Sleep(10 * time.Millisecond)
		err = QueryWithRedis(redisQuery)
	}
}

//KeyUnLock ...
func KeyUnLock(key string) {
	keylock := key + ":lock"
	redisQuery := func(r *redis.Client) error {
		err := r.Del(keylock).Err()
		return err
	}
	QueryWithRedis(redisQuery)
}
