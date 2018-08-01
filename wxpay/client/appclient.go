package client

import (
	"GoPay/wxpay/common"
	"fmt"
	"strings"
	"encoding/json"
	"errors"
	"sort"
	"GoPay/wxpay/config"
	"crypto/md5"
	"encoding/xml"
	"GoPay/wxpay/utils"
	"net/http"
	"io/ioutil"
	"time"
)

type AppClient struct {
	AppId string
	MchId string
}

func NewAppClient(appid, mchid string) *AppClient {
	return &AppClient{
		AppId: appid,
		MchId: mchid,
	}
}

//支付
func (this AppClient) Pay(request common.WxPayAppRequest) (map[string]string, error) {
	request.Appid = this.AppId
	request.MchId = this.MchId
	request.NonceStr = utils.CreateNonceStr()
	request.NotifyUrl = config.NOTIFY_URL
	request.SignType = "MD5"
	request.TradeType = "APP"
	request.SpbillCreateIp = utils.GetIpAddress()
	var response common.WxPayAppResponse
	objByte, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("[gopay->wxpay] json Marshal error, " + err.Error())
	}
	var m = make(map[string]string)
	err = json.Unmarshal(objByte, m)
	if err != nil {
		return nil, errors.New("[gopay->wxpay] json Unmarshal error, " + err.Error())
	}
	m["sign"],err = this.GenerateSign(m)
	if err != nil {
		return nil,err
	}

	xmlStr := utils.MapToXml(m)
	resp,err := http.Post(config.UNIFIEDORDER_URL,"text/xml:charset=UTF-8",strings.NewReader(xmlStr))
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	err = xml.Unmarshal(body,&response)
	if err != nil {
		return nil,errors.New("[gopay->wxpay] xml.Unmarshal error, "+err.Error())
	}
	//组装APP端需要的数据
	var data = make(map[string]string)
	data["appid"] = this.AppId
	data["partnerid"] = this.MchId
	data["prepayid"] = response.PrepayId
	data["package"] = "Sign=WXPay"
	data["noncestr"] = utils.CreateNonceStr()
	data["timestamp"] = fmt.Sprintf("%d",time.Now().Unix())
	data["sign"],err = this.GenerateSign(data)
	return data,nil
}

//支付回调方法
func (this *AppClient) CallBack(w http.ResponseWriter,r http.Request) (common.WxPayNotifyResponse,error)  {
	var returnCode,returnMsg = "FAIL","error"
	defer func() {
	returnStr := "<xml>" +
					"<return_code><![CDATA[%s]]></return_code>" +
					"<return_msg><![CDATA[%s]]></return_msg>" +
				"</xml>"
		returnBody := fmt.Sprintf(returnStr, returnCode, returnMsg)
		w.Write([]byte(returnBody))
	}()
	dataByte,err := ioutil.ReadAll(r.Body)
	var response common.WxPayNotifyResponse
	if err != nil {
		returnMsg = "Body_ERROR"
		return response,errors.New("[gopay->wxpay] callback body error,"+err.Error())
	}
	err = xml.Unmarshal(dataByte,&response)
	if err != nil {
		returnMsg = "解析出错"
		return response,errors.New("[gopay->wxpay] callback xml.Unmarshal error,"+err.Error())
	}
	//判断返回的return_code是否为Success
	if response.ResultCode != "SUCCESS" {
		return response,errors.New("[gopay->wxpay] callback return error: "+response.ReturnMsg)
	}

	//将结构体转为map
	respByte,err := json.Marshal(response)
	if err != nil {
		return response,errors.New("[gopay->wxpay] callback json.Marshal error: "+err.Error())
	}
	var m = make(map[string]string)
	err = json.Unmarshal(respByte,m)
	if err != nil {
		return response,errors.New("[gopay->wxpay] callback json.Unmarshal error: "+err.Error())
	}
	//验签
	signData,err := this.GenerateSign(m)
	if err != nil {
		return response,err
	}
	if signData != m["sign"] {
		return response,errors.New("[gopay->wxpay] callback VerifySign error")
	}
	returnCode,returnMsg = "SUCCESS","OK"
	return response,nil
}

//生成签名
func (this *AppClient) GenerateSign(params map[string]string) (string,error) {
	var data []string
	for k, v := range params {
		if v != "" && k != "sign" {
			data = append(data, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	sort.Strings(data)
	data = append(data, fmt.Sprintf("key=%s", config.MCH_KEY))
	signData := strings.Join(data, "&")
	h := md5.New()
	_,err := h.Write([]byte(signData))
	if err != nil {
		return "",errors.New("[gopay->wxpay] hash write error, "+err.Error())
	}
	result := fmt.Sprintf("%X",h.Sum(nil))
	return result,nil
}


