package main

import (
	"bytes"
	"fmt"
	"regexp"
)

/**
*@Description: regexp 正则表达式
*@Author: imi
*@date: 2019/9/7
 */
func main() {

	matcher, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(matcher)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	reg, err := regexp.Compile("p([a-z]+)ch")
	fmt.Println(reg.MatchString("peeec1h"), err == nil)
	fmt.Println(reg.MatchString("peeech"), err != nil)
	fmt.Println(reg.FindString("peach punch"))
	fmt.Println(reg.FindStringIndex("peach punch"))
	fmt.Println(reg.FindStringSubmatch("peach punch"))
	fmt.Println(reg.FindStringSubmatchIndex("peach punch"))
	fmt.Println(reg.FindAllString("peach punch pinch", 2))

	fmt.Println(reg.FindAllStringSubmatchIndex("peach punch pinch", -1))
	//这个函数提供一个正整数来限制匹配次数
	fmt.Println(reg.FindAllString("peach punch pinch", 2))

	fmt.Println(reg.Match([]byte("peach")))
	//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用于常量。
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	//替换字符串
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	//Func 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
