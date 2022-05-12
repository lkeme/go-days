package main

import (
	"fmt"
	"hello/say"
)

func main() {
	fmt.Println("Hello World!")
	// 引用子包函数
	say.Hello()
}
