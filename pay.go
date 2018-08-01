package main

import (
	payclient "GoPay/alipay/client"
	"GoPay/alipay/config"
	"github.com/golang/glog"
	"GoPay/alipay/common"
	"fmt"
	"net/http"
)

func main()  {
	client,err := payclient.NewPayClient(config.APP_ID,config.SIGN_TYPE,config.PRIVATE_KEY,config.PUBLIC_KEY)
	if err != nil {
		glog.Error(err)
	}
	request := new(common.AliPayTradeAppPayRequest)
	request.Body = "sdfsdf"
	request.Subject = "sdkjfskf"
	result,err := client.Pay(request)
	if err != nil {
		glog.Error(err)
	}
	fmt.Println(request)
	http.HandleFunc("/callBack",func(w http.ResponseWriter,r *http.Request){
		response,err1 := client.CallBack(w,r)
		if err1 != nil {
			glog.Errorln(err1)
		}
	})
}