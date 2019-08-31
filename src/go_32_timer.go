package main

import (
	"fmt"
	"time"
)

/**
*@Description:
我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。Go 的内置 定时器 和 打点器 特性让这些很容易实现。
类似JavaScript中 setTimeout(){}
*@Author: imi
*@date: 2019/8/31
*/
func main() {

	timer1 := time.NewTimer(time.Second)
	<-timer1.C
	fmt.Println("t1 expired ")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("t2 expired")
	}()
	stoped := timer2.Stop()

	if stoped {
		fmt.Println("t2 stopped")
	} else {

	}
}
