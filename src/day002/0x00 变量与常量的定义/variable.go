package main

import "fmt"

func var1() {
	var a int
	var b string
	fmt.Println(a, b)
}

func var2() {
	var a int = 10
	var b string = "abc"
	fmt.Println(a, b)
}

func var3() {
	var a = 10
	var b = "abc"
	fmt.Println(a, b)
}

func var4() {
	var (
		a int
		b string
		c bool
		d float32
	)
	//
	//var a,b,c,d = 10,"abc",true,3.1415926
	//// 可以写在一行
	//
	//var a,b,c,d int = 10,11,12,13
	//// 类型一样的也可以一行同时定义
	//
	fmt.Println(a, b, c, d)
}

func var5() {
	a := 10
	b := "abc"
	fmt.Println(a, b)
}

func main() {
	var1()
	var2()
	var3()
	var4()
	var5()
}
