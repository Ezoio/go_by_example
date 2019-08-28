package main

import (
	"fmt"
	"time"
)

/**
*@Description:
Go 的通道选择器 让你可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性。
*@Author: imi
*@date: 2019/8/27
*/
func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "msg-1"
	}()
	go func() {
		time.Sleep(time.Second * 01)
		c2 <- "msg-2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
