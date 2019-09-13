package main

import (
	"fmt"
	"time"
)

/**
*@Description: 时间格式化
*@Author: imi
*@date: 2019/9/13
 */
func main() {
	p := fmt.Println
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	p(t1.Format("3:04PM"))
	p(t1.Format("Mon Jan _2 15:04:05 2006"))
	p(t1.Format("2006-01-02T15:04:05.999999-07:00"))
	p("------------------")
	p(t1.Format("2019-01-01 06:06:06 "))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	//对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出出时间值得组成。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Printf("%d年%02d月%02d日 %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	//Parse 函数在输入的时间格式不正确是会返回一个错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
