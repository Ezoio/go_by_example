package main

import "fmt"

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
}
