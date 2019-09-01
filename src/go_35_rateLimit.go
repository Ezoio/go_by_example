package main

import (
	"fmt"
	"time"
)

/**
*@Description:
速率限制(英) 是一个重要的控制服务资源利用和质量的途径。
Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
*@Author: imi
*@date: 2019/9/2
*/
func main() {

	request := make(chan int, 5)
	for i := 1; i < 6; i++ {
		request <- i
	}
	close(request)
	limiter := time.NewTicker(time.Millisecond * 100)
	for req := range request {
		<-limiter.C
		fmt.Println("request ", req, time.Now())
	}

	//有时候我们想临时进行速率限制，并且不影响整体的速率控制我们可以通过通道缓冲来实现。
	// 这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i < 4; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 6; i < 11; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	//现在模拟超过 5 个的接入请求。
	// 它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响。
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burstyRequest ", req, time.Now())
	}

}
