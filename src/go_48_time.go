package main

import (
	"fmt"
	"time"
)

/**
*@Description:
Go 对时间和时间段提供了大量的支持；这里是一些例子。
*@Author: imi
*@date: 2019/9/11
*/
func main() {
	p := fmt.Println

	p(time.Now())

	t := time.Date(2020, 6, 6, 5, 6, 6, 6, time.UTC)

	p(t.Year())
	p(t.Month())
	p(t.Day())

	p(t.Hour())
	p(t.Minute())
	p(t.Second())

	p(t.Nanosecond())
	p(t.Location())

	p(t.Weekday())
	p(t.Clock())
	//判断之间前后
	p(t.After(time.Now()))
	p(t.Before(time.Now()))
	p(t.Equal(time.Now()))

	//获取时间差值
	diff := t.Sub(time.Now())
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//移动一个差值的 时间间隔
	p(t.Add(diff))
	p(t.Add(-diff))

	//显示一个数字时钟
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			//cmd := exec.Command("cmd", "/c", "cls")
			//cmd.Stdout = os.Stdout
			//cmd.Run()
			p(time.Now().Clock())
		}
	}()

	select {}

}
