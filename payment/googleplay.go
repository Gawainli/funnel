package payment

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
)

const (
	clientid     = "593458347697-3k289d5gf9m8csoe4o0q8j76pa8bjj9m.apps.googleusercontent.com"
	clientsecret = "eNj3oUFmw2Kscah9hiXFQeeM"
	refreshtoken = "1/YLfsOR24VpPTGFvX1e37KTSn3QdQs577dUJvRvLK8h4"
)

type cacheToken struct {
	accesstoken string
	expiresin   int64
	createtime  int64
}

var mCacheToken cacheToken

//refreshAccessToken refresh access token
func refreshAccessToken() {
	req := httplib.Post("https://accounts.google.com/o/oauth2/token")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(100*time.Second, 30*time.Second)
	req.Debug(true)

	req.Param("grant_type", "refresh_token")
	req.Param("client_id", clientid)
	req.Param("client_secret", clientsecret)
	req.Param("refresh_token", refreshtoken)

	result, err := req.String()

	if err != nil {
		fmt.Println("refresh token error:", err)
	} else {
		fmt.Println("refresh token result:", err)
		var dat map[string]interface{}
		if errRes := json.Unmarshal([]byte(result), &dat); errRes == nil {
			mCacheToken.accesstoken = dat["access_token"].(string)
			mCacheToken.expiresin = int64(dat["expires_in"].(float64))
			mCacheToken.createtime = time.Now().UTC().Unix()
		}
	}
}

//VerifyGooglePay verifiy result for google iap
func VerifyGooglePay(packageName string, productID string, purchaseToken string) (string, error) {
	if mCacheToken.accesstoken != "" {
		cTime := mCacheToken.createtime
		expires := mCacheToken.expiresin
		now := time.Now().UTC().Unix()
		if now > cTime+expires-60 {
			refreshAccessToken()
		}
	} else {
		refreshAccessToken()
	}
	// fmt.Println("access token:", mCacheToken.accesstoken)
	// fmt.Println("token create time:", mCacheToken.createtime)
	// fmt.Println("token expires_in:", mCacheToken.expiresin)
	url := "https://www.googleapis.com/androidpublisher/v2/applications/" + packageName + "/purchases/products/" + productID + "/tokens/" + purchaseToken +
		"?access_token=" + mCacheToken.accesstoken
	fmt.Println("verify url:", url)
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.Debug(true)

	result, err := req.String()

	return result, err
}
