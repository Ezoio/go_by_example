package main

import "fmt"

var c chan string

func Pingpong() {
	i := 1
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From PingPong #%d hi!", i)
		i++
	}

}
func main() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From MAIN #%d hello!", i)
		fmt.Println(<-c)
	}
}
