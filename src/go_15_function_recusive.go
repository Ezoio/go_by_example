package main

import "fmt"

/**
*@Description: 递归
*@Author: imi
*@date: 2019/8/14
 */
func nn(i int) int {
	if i == 0 {
		return 1
	}
	return i * nn(i-1)
}

//func feb(i int) int {
//	if i==0{
//
//	}
//}

func main() {
	fmt.Println(nn(2))
	fmt.Println(nn(3))
	fmt.Println(nn(5))
}
