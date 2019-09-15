package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

/**
*@Description:
SHA1 散列经常用生成二进制文件或者文本块的短标识。
例如，git 版本控制系统大量的使用 SHA1 来标识受版本控制的文件和目录。
这里是 Go中如何进行 SHA1 散列计算的例子。
*@Author: imi
*@date: 2019/9/15
*/
func main() {
	str := "hello go, we  use sha1 hash this String"

	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。
	// 这里我们从一个新的散列开始。
	sha1 := sha1.New()

	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	sha1.Write([]byte(str))

	//这个用来得到最终的散列值的字符切片。
	// Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := sha1.Sum(nil)
	fmt.Println("SHA1》》》》》")
	fmt.Println(bs)
	//运行程序计算散列值并以可读 16 进制格式输出。
	fmt.Printf("%x\n", bs)

	md5 := md5.New()
	md5.Write([]byte(str))
	ms := md5.Sum(nil)
	fmt.Println("MD5》》》》》")
	fmt.Println(ms)
	fmt.Printf("%x", ms)
}
