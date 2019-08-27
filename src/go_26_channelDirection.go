package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func main() {

	pi := make(chan string, 1)
	po := make(chan string, 1)

	ping(pi, "mssssssg")
	pong(pi, po)

	fmt.Println(<-po)

}
