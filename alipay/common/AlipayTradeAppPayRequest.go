package common
//支付请求参数
type AliPayTradeAppPayRequest struct {
	SellerId       string `json:"seller_id,omitempty"`                //收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	Body           string `json:"body,omitempty"`                     //对一笔交易的具体描述信息
	Subject        string `json:"subject,"`                           //商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo     string `json:"out_trade_no"`                       //商户网站唯一订单号
	TimeoutExpress string `json:"timeout_express,omitempty"`          //该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	TimeExpire     string `json:"time_expire,omitempty"`              //绝对超时时间，格式为yyyy-MM-dd HH:mm。
	TotalAmount    string `json:"total_amount"`                       //订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode    string `json:"product_code"`                       //销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY
	GoodsType      string `json:"goods_type,omitempty"`               //商品主类型：0—虚拟类商品，1—实物类商品
	PassbackParams string `json:"passback_params,omitempty"`          //公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	PromoParams    string `json:"promo_params,omitempty"`             //优惠参数
	ExtendParams struct {
		hbFqNum              string `json:"hb_fq_num,omitempty"`               //花呗分期数
		hbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`    //卖家承担收费比例，商家承担手续费传入100，用户承担手续费传入0，仅支持传入100、0两种，其他比例暂不支持
		sysServiceProviderId string `json:"sys_service_provider_id,omitempty"` //系统商编号，该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	} `json:"extend_params,omitempty"`                                //业务扩展参数
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`  //可用渠道，用户只能在指定渠道范围内支付当有多个渠道时用“,”分隔
	DisablePayChannels string `json:"disable_pay_channels,omitempty"` //禁用渠道，用户不可用指定渠道支付当有多个渠道时用“,”分隔
	StoreId            string `json:"store_id,omitempty"`             //商户门店编号。该参数用于请求参数中以区分各门店，非必传项。
}