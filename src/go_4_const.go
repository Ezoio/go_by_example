package main

import (
	"fmt"
	"math"
)
/**
*@Description: 常量 const
1、可以在任何用var地方用const
2、常数表达式可以用于任何精度的数值运算
3、数值型常量通常没有确切的精度，直到被给与一个类型，例如显示类型转换
4、当上下文需要时，一个数可以被给定一个类型
*@Author: imi
*@date: 2019/8/10
*/

const NUM1  = 13434234234234234234 /3E10
const STR_1 = "blabalbalbal"

func main() {

	fmt.Println("NUM1=" ,NUM1)
	fmt.Println("STR_1=" ,STR_1)

	const SIN_VALUE = 343;
	fmt.Println(math.Sin(SIN_VALUE))
}
