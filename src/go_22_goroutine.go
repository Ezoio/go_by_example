package main

import "fmt"

/**
*@Description: 轻量级线程：协程
*@Author: imi
*@date: 2019/8/24
 */
func printNum(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}
func main() {
	//main
	printNum("g1")
	//go启动一个协程
	go printNum("g2")
	//匿名函数启动一个协程
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s, ":", i)
		}
	}("anonymouse function")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
