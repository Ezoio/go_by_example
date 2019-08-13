package main

import "fmt"

/**
*@Description: 初识go函数 function
*@Author: imi
*@date: 2019/8/12
 */
func plus(a int, b int) int {
	return a + b
}
func main() {
	rslt := plus(1, 3)
	fmt.Println(rslt)
}
