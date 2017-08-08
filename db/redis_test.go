package db

import (
	"fmt"
	"testing"

	"github.com/Gawainli/funnel/utils"
	"github.com/go-redis/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func Test_RedisConnect(t *testing.T) {
	err := ConnectRedis("192.168.0.98:6380", "")
	if err != nil {
		t.Error(err.Error())
	}

	s := utils.GenSign("123123123", "213123123")
	fmt.Println(s)
}

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func Test_RedisSet(t *testing.T) {
	Test_RedisConnect(t)
	test := TestStruct{}
	test.Name = "t1"
	test.Age = 10
	test.Addr = "adb"
	js, err := ffjson.Marshal(&test)
	fmt.Println("js:", js)
	fmt.Println("err:", err)

	t2 := TestStruct{}
	err = ffjson.Unmarshal(js, &t2)
	fmt.Println(t2.Age)
	fmt.Println("err:", err)

	query := func(c *redis.Client) error {
		return c.Set("test:id:1", js, 0).Err()
	}
	QueryWithRedis(query)
}

func Test_RedisGet(t *testing.T) {
	Test_RedisConnect(t)
	var value string
	var err error
	query := func(c *redis.Client) error {
		value, err = c.Get("test:id:1").Result()
		return err
	}
	QueryWithRedis(query)
	fmt.Println("value:", value)
	fmt.Println("err:", err)

	t2 := TestStruct{}
	b := []byte(value)
	ffjson.Unmarshal(b, &t2)
	fmt.Println("t2:", t2)
	fmt.Println("t2.name:", t2.Name)
}
