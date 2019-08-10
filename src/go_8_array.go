package main

import "fmt"

/**
*@Description: 数组
*@Author: imi
*@date: 2019/8/10
 */
func main() {
	//声明一个数组，默认值为0 ，元素类型和长度是数组类型的一部分
	var arr []int //
	println(arr)  //[0/0]0x0
	//arr[1] = 1    //不声明长度 不能这样赋值
	arr = []int{1, 2, 3, 4, 5}
	println(arr)          ////[5/5]0xc000041f40
	fmt.Println(arr[1])   //2
	fmt.Println(len(arr)) //5

	var arr1 [] int = [] int{1, 2, 3} //[3/3]0xc000041f08
	println(arr1)

	var arr2 [3][2] int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			arr2[i][j] = i + j
			fmt.Println("arr2", arr2[i][j])
		}
	}
	fmt.Println(arr2)//在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...]

}
