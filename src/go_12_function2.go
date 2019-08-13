package main

import "fmt"

/**
*@Description:
Go 内建多返回值 支持。这个特性在 Go 语言中经常被用到，例如用来同时返回一个函数的结果和错误信息。
*@Author: imi
*@date: 2019/8/12
*/
func vals() (int, int) {
	return 1, 100
}
func vals2() (int, int, string) {
	return 1, 100, "err"
}
func main() {
	a, b := vals()
	fmt.Println("a=", a, " b=", b)

	//使用 _来空白定义符跳过某个参数
	_, c := vals()
	fmt.Println("c=", c)

	fmt.Println(vals())

	fmt.Println(vals2())

	_, _, err := vals2()
	fmt.Println("msg=", err)
}
