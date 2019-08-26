package main

import "fmt"

/**
*@Description:
默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。
可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值。
*@Author: imi
*@date: 2019/8/24
*/
func main() {
	chanVar := make(chan string, 3)
	chanVar <- "123"
	chanVar <- "一二三"
	chanVar <- "aaaaaaaaaaaaa"
	//chanVar <- "@#$#$" //fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-chanVar)
	fmt.Println(<-chanVar)
	fmt.Println(<-chanVar)
	//fmt.Println(<-chanVar)//fatal error: all goroutines are asleep - deadlock!
}
