package utils

import "testing"
import "time"
import "strconv"
import "fmt"

func Test_Sign(t *testing.T) {
	token := GenToken()
	time := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	sign := GenSign(token, time, "adfasdtwet")
	fmt.Println("sign:", sign)
}
