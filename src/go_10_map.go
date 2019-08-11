package main

import "fmt"

/**
*@Description: go内置了map结构
*@Author: imi
*@date: 2019/8/11
 */
func main() {

	m := make(map[string]string) //声明一个map
	m["a"] = "val1"
	m["b"] = "val2"
	m["c"] = "val3"

	fmt.Println(m)
	delete(m, "b")
	fmt.Println(m)

	fmt.Println(len(m))
	b := m["b"]
	fmt.Println(b)

	v := m["a"]
	fmt.Println(v)

	//默认值
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//一行声明并完成初始化
	map1 := map[string]string{"a": "xxx", "b": "bar", "c": "foo"}
	fmt.Print(map1)

}
