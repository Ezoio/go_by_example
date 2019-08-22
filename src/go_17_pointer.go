package main

import "fmt"

func setVal(i int) {
	i = 1
}

func setPtr(p *int) {
	*p = 333333333333333
}

func main() {
	bar := 0
	setVal(bar)
	fmt.Println(bar)

	setPtr(&bar)
	fmt.Println(bar)

	fmt.Println(&bar)

}
