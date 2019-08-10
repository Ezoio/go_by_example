package main

import "fmt"

/**
*@Description: 变量 variable var
*@Author: imi
*@date: 2019/8/10
 */
func main() {
	var a int = 1 // var a = 1
	fmt.Println("a=", a)

	var b string = "string " // var b  = "string "
	fmt.Println("b=", b)

	var c, d float64 = 1.0, 2.0 //var can declare multi variables , var c, d  = 1.0, 2.0
	fmt.Println("c,d", c, d)

	var c1, d1 string //
	fmt.Println("c1,d1", len(c1), d1)

	var c2, d2 int //
	fmt.Println("c1,d1", c2, d2)

	e := "no var variable" //e := "no var variable"
	fmt.Println("e = ", e)

	f := false //var e (string)= "no var variable"  var e = boolean(false);
	fmt.Println("f=", f)

	var f1 = bool(false)
	fmt.Println("f1=", f1)

	var f2 = int(123123)
	fmt.Println("f2=", f2)
	{
		var g = 123123
		fmt.Println("g=", g)
	}
	//fmt.Println("g=",g)//you can't get variable g here
}
