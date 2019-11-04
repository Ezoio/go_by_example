package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("k1", "v111")
	key := os.Getenv("k1")

	fmt.Println(key)
	fmt.Println("all env", os.Environ())
	for _, e := range os.Environ() {
		key := strings.Split(e, "=")
		fmt.Println(key)
	}
}
