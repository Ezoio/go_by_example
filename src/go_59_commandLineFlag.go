package main

import (
	"flag"
	"fmt"
)

/**
*@Description:

命令行标志是命令行程序指定选项的常用方式。
例如，在 wc -l 中，这个 -l 就是一个命令行标志。

Go 提供了一个 flag 包，支持基本的命令行标志解析。
我们将用这个包来实现我们的命令行程序示例。
*@Author: imi
*@date: 2019/10/2
*/

func main() {

	word := flag.String("word", "foo", "a string for argument")
	tag := flag.Bool("tag", false, "a bool value for argument")
	number := flag.Int("number", 0, "a bool value for argument")
	flag.Parse()
	fmt.Println("word", *word)
	fmt.Println("tag", *tag)
	fmt.Println("number", *number)
	fmt.Println("flag.args", flag.Args())

}
