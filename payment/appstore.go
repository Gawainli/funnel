package payment

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
)

//VerifyAppStore verifiy result for appstore iap
func VerifyAppStore(receipt string) (string, error) {
	url := "https://buy.itunes.apple.com/verifyReceipt"

	fmt.Println("url", url)
	fmt.Println("receipt:", receipt)

	req := httplib.Post(url)
	req.SetTimeout(10*time.Second, 10*time.Second)
	// req.Param("receipt-data", receipt)
	req.Debug(true)

	jsonStr := `{"receipt-data":"` + receipt + `"}`
	fmt.Println("jsonStr", jsonStr)
	data := []byte(jsonStr)
	req.Body(data)

	result, err := req.String()

	var dat map[string]interface{}
	json.Unmarshal([]byte(result), &dat)
	if v, ok := dat["status"]; ok {
		status := v.(float64)
		if status == 21007 {
			url1 := "https://sandbox.itunes.apple.com/verifyReceipt"

			fmt.Println("url1", url1)

			req1 := httplib.Post(url1)
			req1.SetTimeout(10*time.Second, 10*time.Second)
			req1.Debug(true)
			req1.Body(data)

			result1, err1 := req1.String()

			return result1, err1
		}
	}

	return result, err
}
