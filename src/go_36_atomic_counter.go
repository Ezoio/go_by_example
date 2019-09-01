package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(2)  指定CPU处理核心数
	var inx uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			//死循环
			for {
				atomic.AddUint64(&inx, 1)
				//让出cpu时间片
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	index := atomic.LoadUint64(&inx)
	fmt.Println("index=", index) //index= 76999234
}
