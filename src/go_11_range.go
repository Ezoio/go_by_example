package main

import "fmt"

/**
*@Description:
range 迭代各种各样的数据结构
*@Author: imi
*@date: 2019/8/11
*/
func main() {

	arr := []int{1, 2, 3, 3, 5, 9}
	// range 迭代数组
	for i, j := range arr {
		//i是数组索引，j是数组的值
		fmt.Print(i, "  ", j)
		if i%2 == 0 {
			fmt.Println(" ", j)
		} else {
			fmt.Println()
		}
	}

	//range 遍历map
	maps := map[string]string{"k1": "v1", "foo": "bar", "k2": "v2"}
	for k, v := range maps {
		fmt.Println(k, " ", v)

	}
	fmt.Println("========================")
	for k1, v1 := range maps {
		fmt.Printf("k=%s,v=%s", k1, v1)
	}

	fmt.Println("========================")
	//遍历字符串
	for m, n := range "0123456789 ABCD abcdefg !@#$%^&*" {
		fmt.Println("m=", m, " n=", n)
	}

}
