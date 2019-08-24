package main

import "fmt"

/**
*@Description: 结构体方法
*@Author: imi
*@date: 2019/8/24
 */
type rect struct {
	width  int
	height int
}

//定义了area方法
func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	r1 := rect{
		width:  3,
		height: 4,
	}
	fmt.Println("area=", r1.area())
}
