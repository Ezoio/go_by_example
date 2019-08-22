package main

import (
	"time"
)

/**
*@Description: 递归
*@Author: imi
*@date: 2019/8/14
 */
func nn(i int) int {
	if i == 0 {
		return 1
	}
	return i * nn(i-1)
}

func feb(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return feb(i-1) + feb(i-2)
	}
}

func main() {
	//fmt.Println(nn(2))
	//fmt.Println(nn(3))
	//fmt.Println(nn(5))
	t1 := time.Now()
	for i := 0; i < 40; i++ {
		print("  ", feb(i))
	}
	t2 := time.Now()
	println()
	print("耗时=", (t2.Sub(t1))/1000000, "ms")
}
