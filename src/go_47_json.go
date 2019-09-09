package main

import (
	"encoding/json"
	"fmt"
)

/**
*@Description:
Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化。
*@Author: imi
*@date: 2019/9/8
*/

//下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type JsonStu struct {
	id       int    `json:"page"`
	name     string `json:"fruits"`
	property []string
}

func main() {
	boolValue, def := json.Marshal(true)
	fmt.Println(string(boolValue), def)

	intValue, _ := json.Marshal(12312)
	fmt.Println(string(intValue))

	floatValue, _ := json.Marshal(12312.343)
	fmt.Println(string(floatValue))

	//这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	mapValue, _ := json.Marshal(map[string]int{"nume": 1, "asd": 3234})
	fmt.Println(string(mapValue))

	stu := &JsonStu{
		id:       0123,
		name:     "11231",
		property: []string{"abc", "123"}}
	stuValue, _ := json.Marshal(stu)
	fmt.Println(string(stuValue))

}
