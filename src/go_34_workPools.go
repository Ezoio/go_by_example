package main

import (
	"fmt"
	"time"
)

/**
*@Description:
在这个例子中，我们将看到如何使用 Go 协程和通道实现一个工作池 。
*@Author: imi
*@date: 2019/9/1
*/

func workers(k int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		fmt.Println("worker", k, "，完成job=", i)
		time.Sleep(time.Second)
		result <- i * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i2 := 0; i2 < 3; i2++ {
		go workers(i2, jobs, result)
	}
	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 9; i++ {
		rlt := <-result
		fmt.Println(rlt)
	}
	//close(result)
}
