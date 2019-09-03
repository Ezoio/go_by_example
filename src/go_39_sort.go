package main

import (
	"fmt"
	"sort"
)

/**
*@Description:
Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。我们首先关注内置数据类型的排序
*@Author: imi
*@date: 2019/9/3
*/
func main() {

	str := []string{"f", "m", "h", "z", "p"}
	sort.Strings(str)
	fmt.Println(str)

	arr := []int{3, 1, 4, 8, 0, 5}
	sort.Ints(arr)
	fmt.Println(arr)

	s := sort.IntsAreSorted(arr)
	fmt.Println(s)

}
