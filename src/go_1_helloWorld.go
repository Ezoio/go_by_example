package main

import "fmt"

/**
 *@Description:
 * $ go run hello-world.go
	 hello world
	 有时候我们想将我们的程序编译成二进制文件。我们可以通过 go build 命来达到目的。

	 $ go build hello-world.go
	 $ ls
	 hello-world	hello-world.go
	 然后我们可以直接运行这个二进制文件。

	 $ ./hello-world
	 hello world
 *@Author: imi
 *@date: 2019/8/10
*/
func main() {
	fmt.Print("hello world!")
}
