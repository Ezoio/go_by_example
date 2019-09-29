package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
*@Description: 行过滤器
*@Author: imi
*@date: 2019/9/29
 */
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := strings.ToUpper(scanner.Text())
		fmt.Println(msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "err:", err)
		os.Exit(1)
	}
}
