package main

import (
	"fmt"
)

func main() {

	message := make(chan string, 1)
	signal := make(chan string, 1)

	select {
	case msg := <-message:
		fmt.Println(msg)
	default:
		fmt.Println("s111111111111111")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println(" s222222 received message   ", msg)
	default:
		fmt.Println("no message")
	}

	select {
	case msg := <-message:
		fmt.Println("s33333 received message  ", msg)
	case sgl := <-signal:
		fmt.Println(" s33333 received signal  ", sgl)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("--------------------------------------------")
	c1, c2 := make(chan int), make(chan string)
	g := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					g <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					g <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1111
	c2 <- "AAAA"

	c1 <- 222
	c2 <- "BBB"
	//close(c1)
	close(c2)
	<-g
}
