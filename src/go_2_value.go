package main

import "fmt"

/**
*@Description: 值
Go 拥有各值类型，包括字符串，整形，浮点型，布尔型等
字符串可以通过 + 连接。
*@Author: imi
*@date: 2019/8/10
*/
func main() {
	fmt.Println("abc" + "bcd")
	fmt.Println(false)
	fmt.Println("7.0/3=", 7.0/3)
	fmt.Println(false || true)
	fmt.Println(!false)
	fmt.Println(!true)
}
