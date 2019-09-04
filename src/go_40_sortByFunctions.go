package main

import (
	"fmt"
	"sort"
)

/**
*@Description:
有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
例如，我们想按照字母的长度而不是首字母顺序对字符串排序。
这里是一个 Go 自定义排序的例子。
*@Author: imi
*@date: 2019/9/3
*/
//使用类型别名
type ByLength []string

//实现sort 接口的三个方法
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	food := []string{"watermelon", "rice", "water", "meat", "sugar"}
	//调用sort.Sort方法使的字符串按照长度大小排序
	sort.Sort(ByLength(food))
	fmt.Println(food)
}
