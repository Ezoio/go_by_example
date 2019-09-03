package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/**
*@Description:
在前面的例子中，我们用互斥锁进行了明确的锁定来让共享的state 跨多个 Go 协程同步访问。
另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。
这个基于通道的方法和 Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，确保每块数据有单独的 Go 协程所有的思路是一致的。
*@Author: imi
*@date: 2019/9/3
*/

//定义读取结构体
type readOps struct {
	key  int
	resp chan int
}

//定义写入结构体
type writeOps struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	reads := make(chan *readOps)
	write := make(chan *writeOps)

	var ops int64 = 0

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-write:
				write.val = state[write.key]
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				reads <- &readOps{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				<-reads
				atomic.AddInt64(&ops, 1)
			}

		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write <- &writeOps{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				<-write
				atomic.AddInt64(&ops, 1)
			}
		}()

	}

	time.Sleep(time.Second)

	finalOps := atomic.LoadInt64(&ops)
	fmt.Println(finalOps)

}

/**
*@Description:  finalOps=2243580
*@Author: imi
*@date: 2019/9/3
 */
