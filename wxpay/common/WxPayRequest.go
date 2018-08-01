package common

type WxPayRequest struct {
	Appid      string `xml:"appid,omitempty",json:"appid,omitempty"`               //应用ID
	MchId      string `xml:"mch_id,omitempty",json:"mch_id,omitempty"`             //商户号
	NonceStr   string `xml:"nonce_str,omitempty",json:"nonce_str,omitempty"`       //随机字符串
	Sign       string `xml:"sign,omitempty",json:"sign,omitempty"`                 //签名
	OutTradeNo string `xml:"out_trade_no,omitempty",json:"out_trade_no,omitempty"` //商户订单号
}
