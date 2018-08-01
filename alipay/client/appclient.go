package client

import (
	"crypto/rsa"
	"fmt"
	"sort"
	"strings"
	"crypto/rand"
	"crypto"
	"encoding/base64"
	"hash"
	"errors"
	"crypto/x509"
	"GoPay/alipay/config"
	"GoPay/alipay/common"
	"time"
	"encoding/json"
	"net/url"
	"net/http"
	"encoding/pem"
)

type AppClient struct {
	AppId         string          //应用ID
	RsaPrivateKey *rsa.PrivateKey //私钥值
	RsaPublicKey  *rsa.PublicKey
	GatewayUrl    string //网关
	Format        string //返回数据格式
	Charset       string //字符集编码
	SignType      string //签名类型
	NotifyUrl     string //回调方法
}

func NewPayClient(appId, signType, privateKey, publicKey string) (*AppClient, error) {
	//加载私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("[gopay->alipay] Sign private key decode error")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("[gopay->alipay] parse private key error, " + err.Error())
	}
	//加载公钥
	block, _ = pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("[gopay->alipay] Sign public key decode error")
	}
	pubKey, err1 := x509.ParsePKCS1PublicKey(block.Bytes)
	if err1 != nil {
		return nil, errors.New("[gopay->alipay] parse public key error, " + err.Error())
	}
	return &AppClient{
		AppId:         appId,
		RsaPrivateKey: priKey,
		RsaPublicKey:  pubKey,
		SignType:      signType,
		NotifyUrl:     config.NOTIFY_URL,
		Charset:       "UTF-8",
		Format:        "json",
	}, nil
}

//支付
func (app *AppClient) Pay(params *common.AliPayTradeAppPayRequest) (string, error) {
	var m = make(map[string]string)
	m["app_id"] = app.AppId
	m["method"] = "alipay.trade.app.pay"
	m["format"] = app.Format
	m["charset"] = app.Charset
	m["sign_type"] = app.SignType
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = app.NotifyUrl
	bizContentByte, err := json.Marshal(params)
	if err != nil {
		return "", errors.New("[gopay->alipay] json Marshal error, " + err.Error())
	}
	m["biz_content"] = string(bizContentByte)
	m["sign"], err = app.GenerateSign(m)
	if err != nil {
		return "", errors.New("[gopay->alipay] generate sign error, " + err.Error())
	}
	var buf []string
	for k, v := range m {
		buf = append(buf, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	return strings.Join(buf, "&"), nil
}

//支付回调
func (app *AppClient) CallBack(w http.ResponseWriter, r *http.Request) (*common.AliPayTradeAppPayResponse, error) {
	var result string
	defer func() {
		w.Write([]byte(result))
	}()
	var m = make(map[string]string)
	var signParams []string
	//解析表单
	r.ParseForm()
	//遍历表单得到返回的参数
	for k, v := range r.Form {
		m[k] = v[0]
		if k == "sign" || k == "sign_type" {
			continue
		}
		signParams = append(signParams, fmt.Sprintf("%s=%s", k, v[0]))
	}
	if m["msg"] != "Success" {
		result = "error"
		return nil, errors.New("[gopay->alipay->callback] Verify Sign error")
	}
	//排序
	sort.Strings(signParams)
	//以&符号连接
	signData := strings.Join(signParams, "&")
	//验签
	err := app.VerifySign(signData, m["sign"])
	if err != nil {
		result = "error"
		return nil, errors.New("[gopay->alipay->callback] Verify Sign error, " + err.Error())
	}
	//将数据转为Json串
	dataByte, err := json.Marshal(m)
	if err != nil {
		result = "error"
		return nil, errors.New("[gopay->alipay->callback] json Marshal error, " + err.Error())
	}
	var response common.AliPayTradeAppPayResponse
	//将json串转为AlipayTradeAppPayResponse
	err = json.Unmarshal(dataByte, &response)
	if err != nil {
		result = "error"
		return nil, errors.New("[gopay->alipay->callback] json Unmarshal error, " + err.Error())
	}
	result = "success"
	return &response, nil
}

//生成签名
func (app *AppClient) GenerateSign(params map[string]string) (string, error) {
	var data []string
	for k, v := range params {
		if v != "" && k != "sign" {
			data = append(data, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	//排序
	sort.Strings(data)
	//以&符号连接
	signData := strings.Join(data, "&")
	var (
		h    hash.Hash
		hash crypto.Hash
	)
	if app.SignType == "RSA2" {
		h = crypto.SHA256.New()
		hash = crypto.SHA256
	} else {
		h = crypto.SHA1.New()
		hash = crypto.SHA1
	}
	_, err := h.Write([]byte(signData))
	if err != nil {
		return "", errors.New("[gopay->alipay] hash write error, " + err.Error())
	}
	digest := h.Sum(nil)
	signByte, err := app.RsaPrivateKey.Sign(rand.Reader, digest, hash)
	if err != nil {
		return "", errors.New("[gopay->alipay] Sign error, " + err.Error())
	}
	signStr := base64.StdEncoding.EncodeToString(signByte)
	return signStr, nil
}

//验证签名
func (app *AppClient) VerifySign(signData, sign string) error {
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return errors.New("[gopay->alipay] sign decode error, " + err.Error())
	}
	var (
		h    hash.Hash
		hash crypto.Hash
	)
	if app.SignType == "RSA2" {
		h = crypto.SHA256.New()
		hash = crypto.SHA256
	} else {
		h = crypto.SHA1.New()
		hash = crypto.SHA1
	}
	_, err = h.Write([]byte(signData))
	if err != nil {
		return errors.New("[gopay->alipay] hash write error, " + err.Error())
	}
	hashByte := h.Sum(nil)
	err = rsa.VerifyPKCS1v15(app.RsaPublicKey, hash, hashByte, signByte)
	if err != nil {
		return errors.New("[gopay->alipay] VerifyPKCS1v15 error, " + err.Error())
	}
	return nil
}
