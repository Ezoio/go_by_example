package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	ZB
)

func main() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
}
