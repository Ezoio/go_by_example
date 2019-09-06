package main

import (
	"fmt"
	s "strings"
)

/**
*@Description: 字符串常见操作函数
*@Author: imi
*@date: 2019/9/6
 */
var print = fmt.Println

func main() {
	print("s.Contains(\"test\",\"s\")=", s.Contains("test", "s"))
	print("s.Compare(\"1a\",\"ssss\")=", s.Compare("1a", "ssss"))
	print("s.ContainsAny(\"asdfasdfasdf\",\"zq\")=", s.ContainsAny("asdfasdfasdf", "zq"))

	print("Contains:  ", s.Contains("test", "es"))
	print("Count:     ", s.Count("test", "t"))
	print("HasPrefix: ", s.HasPrefix("test", "te"))
	print("HasSuffix: ", s.HasSuffix("test", "st"))
	print("Index:     ", s.Index("test", "e"))
	print("Join:      ", s.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", s.Repeat("a", 5))
	print("Replace:   ", s.Replace("foo", "o", "0", -1))
	print("Replace:   ", s.Replace("foo", "o", "0", 1))
	print("Split:     ", s.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", s.ToLower("TEST"))
	print("ToUpper:   ", s.ToUpper("test"))
	print()
	//你可以在 strings包文档中找到更多的函数
	//虽然不是 strings 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制。
	print("Len: ", len("hello"))
	print("Char:", "hello"[1])
}
