package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("服务端错误 :", err)
	}

}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "666 message from golang server， golang  够浪！") //这个写入到w的是输出到客户端的
}
