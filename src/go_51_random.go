package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
*@Description: 随机数
*@Author: imi
*@date: 2019/9/13
 */
func main() {
	//默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Int63n(100))
	fmt.Println(rand.Float32())
	fmt.Println((rand.Uint32() * 5) + 5)
	fmt.Println((rand.Float64() * 5) + 5)

	// 要产生变化的序列，需要给定一个变化的种子。
	// 需要注意的是，如果你出于加密目的，需要使用随机数的话，请使用 crypto/rand 包，此方法不够安全。
	fmt.Println()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Int())
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Int63n(100))
	fmt.Println(r1.Float32())
	fmt.Println((r1.Uint32() * 5) + 5)
	fmt.Println((r1.Float64() * 5) + 5)

	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
