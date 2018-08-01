package common

type WxPayBaseResponse struct {
	ReturnCode string `xml:"return_code" json:"return_code"` //返回状态码
	ReturnMsg  string `xml:"return_msg" json:"return_msg"`   //返回信息
}

type WxPayReturnCodeResponse struct {
	Appid      string `xml:"appid,omitempty"`        //应用APPID
	MchId      string `xml:"mch_id,omitempty"`       //商户号
	DeviceInfo string `xml:"device_info,omitempty"`  //设备号
	NonceStr   string `xml:"nonce_str,omitempty"`    //随机字符串
	Sign       string `xml:"sign,omitempty"`         //签名
	ResultCode string `xml:"result_code,omitempty"`  //业务结果
	ErrCode    string `xml:"err_code,omitempty"`     //错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty"` //错误代码描述
}