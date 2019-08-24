package main

import (
	"fmt"
	"math"
)

/**
*@Description: 接口
*@Author: imi
*@date: 2019/8/24
 */
//几何接口
type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

//实现接口
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//展示方法
func measure(geometry geometry) {
	fmt.Println(geometry)
	fmt.Println("area S=", geometry.area())
	fmt.Println("length L=", geometry.perimeter())
}

func main() {
	r := circle{radius: 9}
	p := rectangle{
		width:  12,
		height: 3.2,
	}
	measure(r)
	measure(p)

}
