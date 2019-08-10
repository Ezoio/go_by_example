package main

import (
	"fmt"
	"time"
)
/**
*@Description:  可以用 switch拓展if-else 写法
*@Author: imi
*@date: 2019/8/10
*/
func main() {
	num := 3
	print("num", num, " is ")
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	 t1 := time.Now()
	switch t1.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("today is weekend")
	default:
		fmt.Println("today is weekday")
	}

	 t2 := time.Now()
	 switch {
	 case 6<t2.Hour()&&t2.Hour()<12:
		 fmt.Println("早上好！")
	 case 12<=t2.Hour()&&t2.Hour()<14:
		 fmt.Println("中午好！")
	 case 14<=t2.Hour()&&t2.Hour()<18:
		 fmt.Println("下午好")

		 
	 }



}
