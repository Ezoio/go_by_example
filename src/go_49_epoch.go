package main

import (
	"fmt"
	"time"
)

/**
*@Description:
一般程序会有获取 Unix 时间的秒数，毫秒数，或者微秒数的需要。
*@Author: imi
*@date: 2019/9/13
*/
func main() {

	now := time.Now()
	sec := now.Second()
	nsec := now.Nanosecond()
	msec := nsec / 1000000

	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(msec)
	fmt.Println(nsec)

	sec1 := now.Unix()
	nsec1 := now.UnixNano()
	msec1 := nsec / 1000000

	fmt.Println()
	fmt.Println(sec1)
	fmt.Println(msec1)
	fmt.Println(nsec1)

	fmt.Println()
	fmt.Println(time.Unix(sec1, 0))
	fmt.Println(time.Unix(0, nsec1))

}
