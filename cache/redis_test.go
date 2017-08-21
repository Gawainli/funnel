package cache

import "testing"

var r *RClient

func Test_RedisConnect(t *testing.T) {
	r = ConnectRedis("192.168.0.98:6381", "")
	if r == nil {
		t.Error("connect redis failed")
	}
}

func Test_RedisPushMulti(t *testing.T) {
	Test_RedisConnect(t)
	s := []int{1, 23, 4235, 5345, 12313, 14536456}
	err := r.RPushMultiInt("test:prmulti", s)
	if err != nil {
		t.Error(err.Error())
	}
}
