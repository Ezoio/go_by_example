package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
*@Description:
Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化。
*@Author: imi
*@date: 2019/9/8
*/

//下面我们将使用这两个结构体来演示自定义类型的编码和解码。
// 属性值首字母必须要大写
type JsonStu struct {
	Id       int
	Name     string
	Property []string
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
		Id:       123,
		Name:     "123",
		Property: []string{"abc", "123"}}
	stuValue, _ := json.Marshal(stu)
	fmt.Println(string(stuValue))

	//jsonStr := `{"name":foo,"bar":{"ord":[1,2,3]}}`
	jsonStr := `{"Id":123,"Name":"123","Property":["abc","123"]}`
	res := &JsonStu{}
	if json.Unmarshal([]byte(jsonStr), res) == nil {
		fmt.Println("res", res)
	}

	//在上面的例子中，我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
	//我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
