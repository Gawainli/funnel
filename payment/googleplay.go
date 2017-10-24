package payment

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
)

const (
	clientid     = "399001718695-htp9eof8mfmtsidp20u0fsgflatd01jj.apps.googleusercontent.com"
	clientsecret = "WBbF3MVDgaSKgzK3OP59dB3L"
	refreshtoken = "1/LH1seSj9h6jUAZdVPDv-uh6DAYvxI_wsIqpHH2ZFcKA"
)

type cacheToken struct {
	accesstoken string
	expiresin   int64
	createtime  int64
}

var mCacheToken cacheToken

//RefreshAccessToken refresh access token
func RefreshAccessToken() (string, error) {
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
		return "", err
	} else {
		fmt.Println("refresh token result:", err)
		var dat map[string]interface{}
		if errRes := json.Unmarshal([]byte(result), &dat); errRes == nil {
			// mCacheToken.accesstoken = dat["access_token"].(string)
			// mCacheToken.expiresin = int64(dat["expires_in"].(float64))
			// mCacheToken.createtime = time.Now().UTC().Unix()
			return dat["access_token"].(string), nil
		} else {
			return "", errRes
		}
	}
}

//VerifyGooglePay verifiy result for google iap
func VerifyGooglePay(packageName string, productID string, purchaseToken string, accessToken string) (string, error) {
	url := "https://www.googleapis.com/androidpublisher/v2/applications/" + packageName + "/purchases/products/" + productID + "/tokens/" + purchaseToken +
		"?access_token=" + accessToken
	fmt.Println("verify url:", url)
	req := httplib.Get(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	req.Debug(true)

	result, err := req.String()

	return result, err
}
