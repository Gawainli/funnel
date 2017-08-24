package nim

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

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
func SetNimReqHeader(req *httplib.BeegoHTTPRequest, appkey string, nonce string, appsecret string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	curTime := time.Now().UTC().Unix()
	checkSum := getCheckSum(appsecret, nonce, strconv.FormatInt(curTime, 10))
	req.Header("AppKey", appkey)
	req.Header("Nonce", nonce)
	req.Header("CurTime", strconv.FormatInt(curTime, 10))
	req.Header("CheckSum", checkSum)
	req.Header("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return nil
}

//BuildSendMsgReq set param for send msg
func BuildSendMsgReq(req *httplib.BeegoHTTPRequest, from string, ope int, to string, msgtype int, body string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("from", from)
	req.Param("ope", strconv.Itoa(ope))
	req.Param("to", to)
	req.Param("type", strconv.Itoa(msgtype))
	req.Param("body", body)
	return nil
}

//BuildKickReq set param for kick
func BuildKickReq(req *httplib.BeegoHTTPRequest, tid string, owner string, member string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("tid", tid)
	req.Param("owner", owner)
	req.Param("member", member)
	return nil
}

//BuildRemoveReq set param for remove
func BuildRemoveReq(req *httplib.BeegoHTTPRequest, tid string, owner string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("tid", tid)
	req.Param("owner", owner)
	return nil
}

//BuildChangeOwnerReq set param for changeOwner
func BuildChangeOwnerReq(req *httplib.BeegoHTTPRequest, tid string, owner string, newowner string, leave int) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("tid", tid)
	req.Param("owner", owner)
	req.Param("newowner", newowner)
	req.Param("leave", strconv.Itoa(leave))
	return nil
}

//BuildAddManagerReq set param for addManager
func BuildAddManagerReq(req *httplib.BeegoHTTPRequest, tid string, owner string, members string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("tid", tid)
	req.Param("owner", owner)
	req.Param("members", members)
	return nil
}

//BuildRemoveManagerReq set param for removeManager
func BuildRemoveManagerReq(req *httplib.BeegoHTTPRequest, tid string, owner string, members string) error {
	if req == nil {
		return errors.New("req is nil")
	}
	req.Param("tid", tid)
	req.Param("owner", owner)
	req.Param("members", members)
	return nil
}
