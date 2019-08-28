package main

import (
	"fmt"
	"time"
)

/**
*@Description:
超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。
得益于通道和 select，在 Go中实现超时操作是简洁而优雅的。
*@Author: imi
*@date: 2019/8/27
*/
func main() {

	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "barrrrrrrr"
	}()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second):
		fmt.Println("超时了")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "foooooooooooo"
	}()
	select {
	case msg1 := <-c2:
		fmt.Println(msg1)
	case msg2 := <-time.After(time.Second * 3):
		fmt.Println(msg2)
	}

}
