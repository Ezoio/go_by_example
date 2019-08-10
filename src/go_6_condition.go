package main

import "fmt"

/**
*@Description: go 里面没有?:的三目运算符，只有if-else 、
即使最简单的判断你还是需要if-else 语句，可以不是用圆括号，但是必须使用花括号
*@Author: imi
*@date: 2019/8/10
*/
func main() {
	if 8/4==3{
		fmt.Println("1")//println & fmt.Println ? 区别
	}
	if a:=1;a==1{
		fmt.Println(a)
	}
	b:=3
	if b<0{
		fmt.Println("负数")
	}else if b == 0{
		fmt.Println("0")
	}else if b >0{
		fmt.Println("正数")
	}
}
