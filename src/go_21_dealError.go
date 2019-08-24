package main

import (
	"errors"
	"fmt"
)

/**
*@Description:
Go 语言使用一个独立的·明确的返回值来传递错误信息的。
这与使用异常的 Java 和 Ruby 以及在 C 语言中经常见到的超重的单返回值/错误值相比，
Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用。
*@Author: imi
*@date: 2019/8/24
*/
type errdev struct {
	arg  int
	prob int
}

func err1(inx int) (int, error) {
	if inx == 33 {
		return -1, errors.New("errors:number=33")
	} else {
		return inx, nil
	}
}

func main() {
	for _, inx := range []int{3, 4, 5, 89, 33, 35} {
		//注意在 if行内的错误检查代码，在 Go 中是一个普遍的用法。
		if r, e := err1(inx); e == nil {
			fmt.Println("program working:", r, e)
		} else {
			fmt.Println("something wrong:", r, e)
		}
	}
}
