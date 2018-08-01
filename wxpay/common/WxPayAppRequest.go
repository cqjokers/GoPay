package common

type WxPayAppRequest struct {
	WxPayRequest
	DeviceInfo     string `xml:"device_info,omitempty",json:"device_info,omitempty"`           //设备号
	SignType       string `xml:"sign_type,omitempty",json:"sign_type,omitempty"`               //签名类型
	Body           string `xml:"body,omitempty",json:"body,omitempty"`                         //商品描述
	Detail         string `xml:"detail,omitempty",json:"detail,omitempty"`                     //商品详情
	Attach         string `xml:"attach,omitempty",json:"attach,omitempty"`                     //附加数据
	FeeType        string `xml:"fee_type,omitempty",json:"fee_type,omitempty"`                 //货币类型
	TotalFee       int32  `xml:"total_fee,omitempty",json:"total_fee,omitempty"`               //总金额
	SpbillCreateIp string `xml:"spbill_create_ip,omitempty",json:"spbill_create_ip,omitempty"` //终端IP
	TimeStart      string `xml:"time_start,omitempty",json:"time_start,omitempty"`             //交易起始时间
	TimeExpire     string `xml:"time_expire,omitempty",json:"time_expire,omitempty"`           //交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty",json:"goods_tag,omitempty"`               //订单优惠标记
	NotifyUrl      string `xml:"notify_url,omitempty",json:"notify_url,omitempty"`             //通知地址
	TradeType      string `xml:"trade_type,omitempty",json:"trade_type,omitempty"`             //交易类型
	LimitPay       string `xml:"limit_pay,omitempty",json:"limit_pay,omitempty"`               //指定支付方式
	SceneInfo      string `xml:"scene_info,omitempty",json:"scene_info,omitempty"`             //场景信息
}
