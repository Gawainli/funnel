package nim

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

func getCheckSum(appSecret string, nonce string, curTime string) string {
	var buf bytes.Buffer
	buf.WriteString(appSecret)
	buf.WriteString(nonce)
	buf.WriteString(curTime)
	h := sha1.New()
	io.WriteString(h, buf.String())
	return fmt.Sprintf("%x", h.Sum(nil))
}

//SetNimReqHeader set nim appkey, nonce, curtime, checkSum
func SetNimReqHeader(req *httplib.BeegoHTTPRequest, appkey string, nonce string, appsecret string) {
	curTime := time.Now().UTC().Unix()
	checkSum := getCheckSum(beego.AppConfig.String("appsecret"), nonce, strconv.FormatInt(curTime, 10))
	req.Header("AppKey", appkey)
	req.Header("Nonce", nonce)
	req.Header("CurTime", strconv.FormatInt(curTime, 10))
	req.Header("CheckSum", checkSum)
	req.Header("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
}
