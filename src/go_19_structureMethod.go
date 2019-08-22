package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	r1 := rect{
		width:  3,
		height: 4,
	}
	fmt.Println(r1.area())
}