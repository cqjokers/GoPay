package common

type TradeFundBill struct {
	FundChannel string  `json:"fund_channel"`       // 交易使用的资金渠道，详见 支付渠道列表(https://doc.open.alipay.com/doc2/detail?treeId=26&articleId=103259&docType=1)
	Amount      float64 `json:"amount,string"`      // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount,string"` // 渠道实际付款金额
}

// 优惠券信息
type VoucherDetail struct {
	Id                 string  `json:"id"`                         // 券id
	Name               string  `json:"name"`                       // 券名称
	Type               string  `json:"type"`                       // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券 ALIPAY_DISCOUNT_VOUCHER - 折扣券 ALIPAY_ITEM_VOUCHER - 单品优惠 注：不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码
	Amount             float64 `json:"amount,string"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute float64 `json:"merchant_contribute,string"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    float64 `json:"other_contribute,string"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string  `json:"memo"`                       // 优惠券备注信息
}

//APP支付响应
type AliPayTradeAppPayResponse struct {
	AliPayResponse
	NotifyTime        string          `json:"notify_time"`         //通知时间
	NotifyType        string          `json:"notify_type"`         //通知类型
	NotifyId          string          `json:"notify_id"`           //通知校验ID
	AppId             string          `json:"app_id"`              //支付宝分配给开发者的应用Id
	Charset           string          `json:"charset"`             //编码格式
	Version           string          `json:"version"`             //接口版本
	SignType          string          `json:"sign_type"`           //签名类型
	TradeNo           string          `json:"trade_no"`            //支付宝交易号
	OutTradeNo        string          `json:"out_trade_no"`        //商户订单号
	OutBizNo          string          `json:"out_biz_no"`          //商户业务号
	BuyerId           string          `json:"buyer_id"`            //买家支付宝用户号
	BuyerLogonId      string          `json:"buyer_logon_id"`      //买家支付宝账号
	SellerId          string          `json:"seller_id"`           //卖家支付宝用户号
	SellerEmail       string          `json:"seller_email"`        //卖家支付宝账号
	TradeStatus       string          `json:"trade_status"`        //交易状态
	TotalAmount       float64         `json:"total_amount"`        // 订单金额
	ReceiptAmount     float64         `json:"receipt_amount"`      //实收金额
	InvoiceAmount     float64         `json:"invoice_amount"`      //开票金额
	BuyerPayAmount    float64         `json:"buyer_pay_amount"`    //付款金额
	PointAmount       float64         `json:"point_amount"`        //集分宝金额
	RefundFee         float64         `json:"refund_fee"`          //总退款金额
	Subject           string          `json:"subject"`             //订单标题
	Body              string          `json:"body"`                //商品描述
	PassbackParams    string          `json:"passback_params"`     //回传参数
	GmtCreate         string          `json:"gmt_create"`          //交易创建时间
	GmtPayment        string          `json:"gmt_payment"`         //交易付款时间
	GmtRefund         string          `json:"gmt_refund"`          //交易退款时间
	GmtClose          string          `json:"gmt_close"`           //交易结束时间
	FundBillList      []TradeFundBill `json:"fund_bill_list"`      // 交易支付使用的资金渠道
	VoucherDetailList []VoucherDetail `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
}
