package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//GenToken 创建token
func GenToken() string {
	return bson.NewObjectId().Hex()
}

//GenSign ...
func GenSign(token string, timestamp string, signKey string) string {
	data := []byte(token + timestamp + signKey)
	buf := md5.Sum(data)
	s := hex.EncodeToString(buf[:])
	return strings.ToUpper(s)
}
