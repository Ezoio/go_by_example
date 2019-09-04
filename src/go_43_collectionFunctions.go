package main

/**
*@Description:
我们经常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，
或者将所有的项通过一个自定义函数映射到一个新的集合上。
在某些语言中，会习惯使用泛型。Go 不支持泛型，在 Go 中，
当你的程序或者数据类型需要时，通常是通过组合的方式来提供操作函数。
*@Author: imi
*@date: 2019/9/4
*/
func indexStr(str []string, s string) int {
	for i, v := range str {
		if s == v {
			return i
		}
	}
	return -1
}

func includeIndexStr(str []string, s string) bool {
	return indexStr(str, s) > 0
}

func main() {

}
