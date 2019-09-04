package main

import "os"

func main() {

	panic("something is error")

	_, err := os.Create("temp/panic.txt")
	if err != nil {
		panic(err)
	}
}
