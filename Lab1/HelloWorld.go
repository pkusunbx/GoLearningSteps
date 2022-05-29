package main

import (
	"fmt"
	"strconv"
)

var (
	x int
	y bool
)

const zz = "const value"

//不带声明格式的变量只能在函数体中出现
//g, h := 123, "hello"

var ee, ff = 123, "hello"

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

	fmt.Println(ee, ff)

	fmt.Println(&ee, &ff)

	//cc := "ccccccc"

	xx, yy := 1, 2
	xx, yy = yy, xx
	fmt.Println(xx, yy)

	fmt.Println(zz)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	fmt.Println()

	const (
		v1 = iota
		v2
		v3
		v4 = "ha"
		v5
		v6 = 100
		v7
		v8 = iota
		v9
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9)

	const (
		a1 = 1 << iota
		a2 = 3 << iota
		a3
		a4
	)
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)
	fmt.Println("a3=", a3)
	fmt.Println("a4=", a4)

	var a5 bool = true
	var a6 bool = false
	if a5 && a6 {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a5 || a6 {
		fmt.Printf("第二行 - 条件为 true\n")
	}

	var a7 int = 4
	var ptr *int
	ptr = &a7
	fmt.Println(ptr)
	fmt.Println(*ptr)

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var a8 int = 100
	var a9 int = 200
	ret := max(a8, a9)
	fmt.Printf("最大值是：%d\n", ret)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
