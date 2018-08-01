package common

type WxPayQueryResponse struct {
	WxPayBaseResponse
	WxPayReturnCodeResponse
	Openid             string `xml:"openid,omitempty"`               //用户标识
	IsSubscribe        string `xml:"is_subscribe,omitempty"`         //是否关注公众账号
	TradeType          string `xml:"trade_type,omitempty"`           //交易类型
	TradeState         string `xml:"trade_state,omitempty"`          //交易状态
	BankType           string `xml:"bank_type,omitempty"`            //付款银行
	TotalFee           int32  `xml:"total_fee,omitempty"`            //总金额
	FeeType            string `xml:"fee_type,omitempty"`             //货币种类
	CashFee            int32  `xml:"cash_fee,omitempty"`             //现金支付金额
	CashFeeType        string `xml:"cash_fee_type,omitempty"`        //现金支付货币类型
	SettlementTotalFee int32  `xml:"settlement_total_fee,omitempty"` //应结订单金额
	CouponFee          int32  `xml:"coupon_fee,omitempty"`           //代金券金额
	CouponCount        int32  `xml:"coupon_count,omitempty"`         //代金券使用数量
	TransactionId      string `xml:"transaction_id,omitempty"`       //微信支付订单号
	OutTradeNo         string `xml:"out_trade_no,omitempty"`         //商户订单号
	Attach             string `xml:"attach,omitempty"`               //附加数据
	TimeEnd            string `xml:"time_end,omitempty"`             //支付完成时间
	TradeStateDes      string `xml:"trade_state_des,omitempty"`      //交易状态描述
}
