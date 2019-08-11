package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	println(s)
	fmt.Println(s)
	fmt.Println(s[1])

	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)

	arr := make([]string, len(s))
	fmt.Println(arr)
	copy(arr, s)

	fmt.Println("arr=", arr)

	arr1 := arr[1:3] //不包含arr[2]
	fmt.Println("arr1=", arr1)

	arr2 := arr[:5] //包含arr[4]
	fmt.Println("arr2=", arr2)

	arrays := make([][]int, 5)
	fmt.Println("before arrays=", arrays)
	for i := 0; i < 5; i++ {
		//Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
		arrays[i] = make([]int, 5)
		//arrays[i] = make([]int,i+1)
		for j := 0; j < i+1; j++ {
			arrays[i][j] = i + j
		}
	}
	fmt.Println("after arrays=", arrays)

}
