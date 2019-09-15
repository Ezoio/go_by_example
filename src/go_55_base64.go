package main

import (
	"encoding/base64"
	"fmt"
)

/**
*@Description:
Go 提供内建的 base64 编解码支持。
*@Author: imi
*@date: 2019/9/15
*/
func main() {
	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	str := "QWERTYUIOP{(*&^%$#@!:/.,;"
	std64EnStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("标准64编码", std64EnStr)
	std64DeStr, err := base64.StdEncoding.DecodeString(std64EnStr)
	fmt.Println("标准64解码", std64DeStr, err)

	//标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。
	str1 := "123"
	url64EnStr := base64.URLEncoding.EncodeToString([]byte(str1))
	fmt.Println("标准64编码", url64EnStr)
	url64DeStr, err := base64.StdEncoding.DecodeString(url64EnStr)
	fmt.Println("标准64解码", url64DeStr, err)
}
