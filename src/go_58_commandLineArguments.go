package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(args)
	fmt.Println(argsWithoutProg)

}
