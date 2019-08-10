package main

import "fmt"
/**
*@Description: for循环 单循环，经典循环，无限循环（bread return 终止循环）
*@Author: imi
*@date: 2019/8/10
*/
func main() {
	i := 1
	for i<3{
		fmt.Print(" ",i)
		i =i+1//i++
	}
	fmt.Println()

	for i:=3;i<9;i++{
		fmt.Print(" ",i)
	}
	fmt.Println()
	for {
		fmt.Println("endless loop")
		//break
		return
	}
}
