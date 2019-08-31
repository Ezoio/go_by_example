package main

import (
	"fmt"
	"time"
)

/**
*@Description:
关闭 一个通道意味着不能再向这个通道发送值了。
这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
*@Author: imi
*@date: 2019/8/31
*/
func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs //通道关闭后 more = false
			if more {
				fmt.Println(" received j=", j, "more = ", more)
			} else {
				fmt.Println(" all jobs done ", j, "more = ", more)
				done <- true
				return
			}
		}
	}()

	for i := 3; i > 0; i-- {
		jobs <- i
		fmt.Println("  jobs send ", i)
		time.Sleep(time.Second)
		fmt.Println(done)
	}
	fmt.Println(" all  jobs send ")
	close(jobs)
	fmt.Println(done)
	<-done
}
