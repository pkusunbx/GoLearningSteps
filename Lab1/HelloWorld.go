package main

import (
	"fmt"
	"strconv"
)

var (
	x int
	y bool
)

func main() {
	fmt.Printf("Hello" + "," + "World!\n")
	var age = 18
	a := strconv.Itoa(age)
	fmt.Println("age: " + a + "\n")
	var stockcode = 123
	var enddate = "2022-05-22"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var b bool = true
	if b != false {
		fmt.Println("b != false")
	}

	var c string = "runoob"
	fmt.Println(c)

	//指定变量类型，如果没有初始化，则变量默认为零值。零值就是变量没有做初始化时系统默认设置的值。
	var d, e int = 1, 2
	fmt.Println(d, e)

	var f int
	fmt.Println(f)

	var g bool
	fmt.Println(g)

	var i_ int
	var f_ float64
	var b_ bool
	var s_ string
	fmt.Printf("%v %v %v %q\n", i_, f_, b_, s_)

	h := "ahhahaaha"
	fmt.Println(h, x, y)
}
