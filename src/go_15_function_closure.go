package main

import "fmt"

/**
*@Description: 闭包
*@Author: imi
*@date: 2019/8/14
 */
func add1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	addFunc := add1()
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	addFunc1 := add1()
	fmt.Println(addFunc1())
}
