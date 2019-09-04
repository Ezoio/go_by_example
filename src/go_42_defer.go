package main

import (
	"errors"
	"fmt"
	"os"
)

/**
*@Description:
Defer 被用来确保一个函数调用在程序执行结束前执行。
同样用来执行一些清理工作。
defer 用在像其他语言中的ensure 和 finally用到的地方。
*@Author: imi
*@date: 2019/9/4
*/
var p = fmt.Println

func main() {

	f := createFile("test1")
	defer closeFile(f)
	writeFile(f)

}
func createFile(fileName string) *os.File {
	p("create file")
	f, err := os.Create(fileName)
	if err != nil {
		panic(errors.New("create file failure"))
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("write something ")
}

func closeFile(f *os.File) {
	fmt.Println("close file")
}
