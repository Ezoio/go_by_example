package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx")
	os.Exit(3)
}
