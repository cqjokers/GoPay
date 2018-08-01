package utils

import (
	"math/rand"
	"time"
	"strings"
	"net"
	"bytes"
	"fmt"
)

//生成随机字符串
func CreateNonceStr() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	var str []string
	for i := 0; i < 32; i++ {
		index := rand.Intn(31)
		str = append(str, chars[index:index+1])
	}
	return strings.Join(str, "")
}

//获取IP地址
func GetIpAddress() string  {
	addr,_ := net.InterfaceAddrs()
	for _,address := range addr{
		if ipNet,ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback(){
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

//map转XML,另一种方式是使用自带的xml包生成，感觉使用上有点麻烦所以未使用
/*
	type CDATA struct {
		Text string `xml:",cdata"`
	}

	type Person struct {
		XMLName   xml.Name `xml:"xml"`
		Id        int      `xml:"id"`
		FirstName string   `xml:"name>first,omitempty"`
		LastName  CDATA   `xml:"name>last"`
		Age       CDATA
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}
	v := &Person{Id: 13, FirstName: "John", LastName: CDATA{"Doe"}, Age: CDATA{"25"}}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
	结果如下：
	<xml>
      <id>13</id>
      <name>
          <first>John</first>
          <last><![CDATA[Doe]]></last>
      </name>
      <Age><![CDATA[25]]></Age>
      <Married>false</Married>
      <City>Hanga Roa</City>
      <State>Easter Island</State>
  </xml>
*/
func MapToXml(m map[string]string) string{
	bufStr := bytes.NewBufferString("")
	for k,v := range m{
		bufStr.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>",k,v,k))
	}
	return fmt.Sprintf("<xml>%s</xml>",bufStr.String())
}
