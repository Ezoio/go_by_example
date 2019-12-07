package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
*@Description:
通道 是连接多个 Go 协程的管道。你可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收。
*@Author: imi
*@date: 2019/8/24
*/
func main() {
	//msg := make(chan string)
	//flag := make(chan bool)
	//
	//go func() {
	//	msg <- "this message from a goroutine "
	//	flag <- true
	//}()
	//
	////m :=msg
	////mm <- msg
	//mm := <-msg
	//fmt.Println(mm)
	//for v := range flag{
	//	fmt.Println(v)
	//}

	//c := make(chan bool)//make(chan bool,1)是异步不阻塞的
	//go func() {
	//	fmt.Println("go go go")
	//	<-c
	//}()
	//c <- true
	test4CPUS()
}

func test4CPUS() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool,10)
	//c := make(chan bool,10)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counters(&wg, i)
	}
	wg.Wait()
	//for i := 0; i < 10; i++ {
	//	<-c
	//}
}

func Counters(c *sync.WaitGroup, idx int) {
	a := 1
	for i := 0; i < 99999999; i++ {
		a += i
	}
	fmt.Println(idx, a)
	c.Done()
}

//10个goroutine 打印输出 问题
//一种chan 10个缓存
//第二种是同步包
