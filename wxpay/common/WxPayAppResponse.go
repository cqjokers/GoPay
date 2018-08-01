package common

type WxPayAppResponse struct {
	WxPayBaseResponse
	WxPayReturnCodeResponse
	TradeType string `xml:"trade_type,omitempty"` //交易类型
	PrepayId  string `xml:"prepay_id,omitempty"`  //预支付交易会话标识
}
