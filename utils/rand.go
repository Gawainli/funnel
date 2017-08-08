package utils

import (
	"math/rand"
	"time"
)

func GenRandIntN(e int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(e)

}

//GenRandIntRange 随机区间[s ,e)
func GenRandIntRange(s int, e int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(e-s) + s

}
