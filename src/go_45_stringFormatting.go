package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

/**
*@Description: Go 在传统的printf 中对字符串格式化提供了优异的支持。这里是一些基本的字符串格式化的人物的例子。
*@Author: imi
*@date: 2019/9/7
 */
func main() {
	p := point{10, 5}
	//
	fmt.Printf("%v\n", p)
	//%+v 的格式化输出内容将包括结构体
	fmt.Printf("%+v\n", p)
	//形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p)
	//值的类型
	fmt.Printf("%T\n", p)
	//格式化布尔值
	fmt.Printf("%t\n", false)
	//整数格式化
	fmt.Printf("%d \n", 1234)
	//二进制
	fmt.Printf("%b \n", 1234)
	//16进制
	fmt.Printf("%x \n", 1234)
	//这个输出给定整数的对应字符
	fmt.Printf("%c \n", 65)
	//%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	fmt.Printf("%e \n", 913.999)
	fmt.Printf("%E \n", 65333333.234343434343434)
	//%s 基本的字符串输出
	fmt.Printf("%s \n", "\"golang\"")
	//%q 像golang源码输入
	fmt.Printf("%q \n", "\"golang\"")
	//输出base16
	fmt.Printf("%x \n", "A")
	//输出指针%p
	fmt.Printf("%p \n", &p)

	fmt.Println("=================默认右对齐")
	fmt.Printf("|%6d|%6d|\n", 13, 3434)
	//浮点数使用 宽度.精度 来格式化
	fmt.Printf("|%6.2f|%6.2f|\n", 12.2, 399.344)
	fmt.Printf("|%6.2f|%6.2f|\n", 12.2, 399.345)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	//使用- 左对齐
	fmt.Println("=================左对齐")
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.2, 399.344)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.2, 399.345)
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	//到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。
	//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
