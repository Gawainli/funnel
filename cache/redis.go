package cache

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// var client *redis.Client

//RClient redis client 拓展
type RClient struct {
	*redis.Client
}

//RClient redis client 拓展
// type RClient redis.Client

//ConnectRedis ...
func ConnectRedis(addr string, pwd string) (r *RClient) {
	r = &RClient{}
	options := &redis.Options{}
	options.Addr = addr
	options.Password = pwd
	r.Client = redis.NewClient(options)
	pong, err := r.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		r = nil
	}
	return
}

//KeyLock SetNx实现的分布式锁
func (r *RClient) KeyLock(key string) {
	keylock := key + "@lock"
	err := r.SetNX(keylock, strconv.FormatInt(time.Now().Unix(), 10), 5*time.Second).Err()
	for err != nil {
		time.Sleep(10 * time.Millisecond)
	}
}

//KeyUnLock ...
func (r *RClient) KeyUnLock(key string) {
	keylock := key + "@lock"
	r.Del(keylock).Err()
}

//LRangeInt ...
func (r *RClient) LRangeInt(key string, s int, e int) ([]int, error) {
	var strs []string
	var val []int
	var err error
	strs, err = r.LRange(key, int64(s), int64(e)).Result()
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

//LRangeString ...
func (r *RClient) LRangeString(key string, s int, e int) ([]string, error) {
	var strs []string
	var err error
	strs, err = r.LRange(key, int64(s), int64(e)).Result()
	return strs, err
}

//GetIncrKey ...
func (r *RClient) GetIncrKey(key string, incr int) (int64, error) {
	val, err := r.IncrBy(key, int64(incr)).Result()
	return val, err
}

//SAddMultiString ...
func (r *RClient) SAddMultiString(key string, strs []string) error {
	inf := []interface{}{}

	for _, v := range strs {
		inf = append(inf, v)
	}

	return r.SAdd(key, inf...).Err()
}

//SAddMultiInt ...
func (r *RClient) SAddMultiInt(key string, vals []int) error {
	inf := []interface{}{}

	for _, v := range vals {
		inf = append(inf, v)
	}

	return r.SAdd(key, inf...).Err()
}

//RPushMultiInt ...
func (r *RClient) RPushMultiInt(key string, val []int) error {
	inf := []interface{}{}
	for _, v := range val {
		inf = append(inf, v)
	}
	return r.RPush(key, inf...).Err()
}

//RPushMultiBytes ...
func (r *RClient) RPushMultiBytes(key string, val [][]byte) error {
	inf := []interface{}{}
	for _, v := range val {
		inf = append(inf, v)
	}
	return r.RPush(key, inf...).Err()
}
