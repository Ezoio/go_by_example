package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 300)
	i := 0
	go func() {
		for t := range ticker.C {
			i++
			fmt.Println("ticker at :", t, " i=", i)
		}
	}()

	time.Sleep(time.Millisecond * 1000)
	ticker.Stop()
	fmt.Println("ticker stopped")

}
