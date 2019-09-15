package main

import (
	"fmt"
	"net/url"
)

/**
*@Description:
URL 提供了一个统一资源定位方式。这里了解一下 Go 中是如何解析 URL 的。
*@Author: imi
*@date: 2019/9/15
*/
func main() {

	url1 := "postgres://user:pass@host.com:5432/path?k=v#f"
	url2 := "https://vip.tmall.com/vip/index.htm?spm=875.7931836/B.2016004.3.66144265IxdkLq&acm=lb-zebra-148799-667861." +
		"1003.4.2429983&scm=1003.4.lb-zebra-148799-667861.OTHER_14561845383563_2429983"

	u1, e1 := url.Parse(url1)
	u2, e2 := url.Parse(url2)

	fmt.Println(u1, e1)
	fmt.Println("ForceQuery=", u1.ForceQuery)
	fmt.Println("Fragment=", u1.Fragment)
	fmt.Println("Host=", u1.Host)
	fmt.Println("Qpaque=", u1.Opaque)
	fmt.Println("Path=", u1.Path)
	fmt.Println("RawPath=", u1.RawPath)
	fmt.Println("Scheme=", u1.Scheme)
	fmt.Println("User=", u1.User)
	fmt.Println("EscapedPath=", u1.EscapedPath())
	fmt.Println("Hostname=", u1.Hostname())
	fmt.Println("IsAbs=", u1.IsAbs())
	fmt.Print("MarshalBinary=")
	fmt.Println(u1.MarshalBinary())
	fmt.Println("Port=", u1.Port())
	fmt.Println("RequestURI=", u1.RequestURI())
	fmt.Println("String=", u1.String())
	fmt.Println(u1.RawQuery)
	m, _ := url.ParseQuery(u1.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	fmt.Println("\n========================================================" +
		"========================================================\n\n")
	fmt.Println(u2, e2)
	fmt.Println("ForceQuery=", u2.ForceQuery)
	fmt.Println("Fragment=", u2.Fragment)
	fmt.Println("Host=", u2.Host)
	fmt.Println("Qpaque=", u2.Opaque)
	fmt.Println("Path=", u2.Path)
	fmt.Println("RawPath=", u2.RawPath)
	fmt.Println("Scheme=", u2.Scheme)
	fmt.Println("User=", u2.User)
	fmt.Println("EscapedPath=", u2.EscapedPath())
	fmt.Println("Hostname=", u2.Hostname())
	fmt.Println("IsAbs=", u2.IsAbs())
	fmt.Print("MarshalBinary=")
	fmt.Println(u2.MarshalBinary())
	fmt.Println("Port=", u2.Port())
	fmt.Println("RequestURI=", u2.RequestURI())
	fmt.Println("String=", u2.String())
	m1, _ := url.ParseQuery(u2.RawQuery)
	fmt.Println(m1)
	fmt.Println(m1["acm"][0])
}
