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
	url3 := "https://www.tmall.com/?spm=3700.7045653.a2226mz.1.JcVpSG"

	u1, e1 := url.Parse(url1)
	u2, e2 := url.Parse(url2)
	u3, e3 := url.Parse(url3)

	fmt.Println(u1, e1)
	fmt.Println(u2, e2)
	fmt.Println(u3, e3)

}
