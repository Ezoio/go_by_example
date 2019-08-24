package main

import "fmt"

/**
*@Description:
通道 是连接多个 Go 协程的管道。你可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收。
*@Author: imi
*@date: 2019/8/24
*/
func main() {
	msg := make(chan string)

	go func() {
		msg <- "this message from a goroutine "
	}()

	//m :=msg
	//mm <- msg
	mm := <-msg
	fmt.Println(mm)
}
