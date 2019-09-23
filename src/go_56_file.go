package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
*@Description:
读写文件在很多程序中都是必须的基本任务。
*@Author: imi
*@date: 2019/9/23
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(fileName string) bool {
	var isExist = true
	if _, err := os.Stat(fileName); os.IsExist(err) {
		isExist = false
	}
	return isExist
}
func main() {

	//读文件
	data1, err := ioutil.ReadFile("abc")
	fmt.Println(string(data1), err)

	//控制文件读取流程
	file1, err := os.Open("abc")
	check(err)
	fmt.Println(&file1, err)

	//控制文件读取一些字节
	b1 := make([]byte, 5)
	n1, err := file1.Read(b1)
	check(err)
	fmt.Println(n1, b1, string(b1), err)

	//从一个文件的某个位置读取
	n2, err := file1.Seek(6, 0)
	b2 := make([]byte, 2)
	check(err)
	n3, err := file1.Read(b2)
	fmt.Println(n2, b2, n3, string(b2), err)

	n4, err := file1.Seek(4, 0)
	b3 := make([]byte, 4)
	check(err)
	n5, err := io.ReadAtLeast(file1, b3, 2)
	fmt.Println(n4, b3, n5, string(b3), err)

	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = file1.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(file1)
	b3, err = r4.Peek(4)
	check(err)
	fmt.Println(string(b3))

	defer file1.Close()

	//写文件
	d1 := []byte("runner ,just do it !")
	ioutil.WriteFile("abc", d1, 0644)

}
