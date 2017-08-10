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
	keylock := key + "@lock"
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
	keylock := key + "@lock"
	redisQuery := func(r *redis.Client) error {
		err := r.Del(keylock).Err()
		return err
	}
	QueryWithRedis(redisQuery)
}

func SetKey(key string, buf []byte, ex time.Duration) error {
	setQuery := func(r *redis.Client) error {
		err := r.Set(key, buf, ex).Err()
		return err
	}
	return QueryWithRedis(setQuery)
}

func SetKeyString(key string, val string, ex time.Duration) error {
	setQuery := func(r *redis.Client) error {
		err := r.Set(key, val, ex).Err()
		return err
	}
	return QueryWithRedis(setQuery)
}

func GetKeyString(key string) (str string, err error) {
	getQuery := func(r *redis.Client) error {
		str, err = r.Get(key).Result()
		return err
	}
	err = QueryWithRedis(getQuery)
	if err != nil {
		return str, err
	}
	return
}

func GetKey(key string) (buf []byte, err error) {
	getQuery := func(r *redis.Client) error {
		buf, err = r.Get(key).Bytes()
		return err
	}
	err = QueryWithRedis(getQuery)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func ExistsKey(key string) int {
	re := -1
	existsQuery := func(r *redis.Client) error {
		i64, err := r.Exists(key).Result()
		re = int(i64)
		return err
	}
	QueryWithRedis(existsQuery)
	return re
}

func DelKey(key string) {
	delQuery := func(r *redis.Client) error {
		err := r.Del(key).Err()
		return err
	}
	QueryWithRedis(delQuery)
}

func SAddMultiString(key string, strs []string) error {
	inf := []interface{}{}

	for _, v := range strs {
		inf = append(inf, v)
	}

	query := func(c *redis.Client) error {
		err := c.SAdd(key, inf...).Err()
		return err
	}
	return QueryWithRedis(query)
}

func SAddMultiInt(key string, vals []int) error {
	inf := []interface{}{}

	for _, v := range vals {
		inf = append(inf, v)
	}

	query := func(c *redis.Client) error {
		err := c.SAdd(key, inf...).Err()
		return err
	}
	return QueryWithRedis(query)
}

func RPushMultiInt(key string, val []int) error {
	inf := []interface{}{}
	for _, v := range val {
		inf = append(inf, v)
	}
	query := func(c *redis.Client) error {
		err := c.RPush(key, inf...).Err()
		return err
	}
	return QueryWithRedis(query)
}

func RPushMultiBytes(key string, val [][]byte) error {
	inf := []interface{}{}
	for _, v := range val {
		inf = append(inf, v)
	}
	query := func(c *redis.Client) error {
		err := c.RPush(key, inf...).Err()
		return err
	}
	return QueryWithRedis(query)
}

func LRangeString(key string, s int, e int) ([]string, error) {
	var strs []string
	var err error
	query := func(c *redis.Client) error {
		strs, err = c.LRange(key, int64(s), int64(e)).Result()
		return err
	}
	QueryWithRedis(query)
	return strs, err
}

func LRangeInt(key string, s int, e int) ([]int, error) {
	var strs []string
	var val []int
	var err error
	query := func(c *redis.Client) error {
		strs, err = c.LRange(key, int64(s), int64(e)).Result()
		return err
	}
	QueryWithRedis(query)

	val = make([]int, len(strs))
	for i, v := range strs {
		k, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		val[i] = k
	}
	return val, err
}

func GetIncrKey(key string, incr int) (int64, error) {
	var val int64
	var err error
	query := func(c *redis.Client) error {
		val, err = c.IncrBy(key, int64(incr)).Result()
		return err
	}
	err = QueryWithRedis(query)
	return val, err
}

func SMembers(key string) ([]string, error) {
	val := []string{}
	var err error
	redisQuery := func(r *redis.Client) error {
		val, err = r.SMembers(key).Result()
		return err
	}
	err = QueryWithRedis(redisQuery)
	return val, err
}

func LSetString(key string, idx int64, val string) error {
	redisQuery := func(r *redis.Client) error {
		err := r.LSet(key, idx, val).Err()
		return err
	}
	return QueryWithRedis(redisQuery)
}

func LSetInt(key string, idx int64, val int) error {
	redisQuery := func(r *redis.Client) error {
		err := r.LSet(key, idx, val).Err()
		return err
	}
	return QueryWithRedis(redisQuery)
}
