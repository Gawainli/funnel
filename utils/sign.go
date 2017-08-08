package utils

import (
	"crypto/md5"
	"encoding/hex"

	"gopkg.in/mgo.v2/bson"
)

const signKey = "funnel"

//GenToken 创建token
func GenToken() string {
	return bson.NewObjectId().Hex()
}

//GenSign ...
func GenSign(token string, timestamp string) string {
	data := []byte(token + timestamp + signKey)
	buf := md5.Sum(data)
	s := hex.EncodeToString(buf[:])
	return s
}
