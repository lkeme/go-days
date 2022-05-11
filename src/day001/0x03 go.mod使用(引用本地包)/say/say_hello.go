package say

import "fmt"

func Hello() {
	fmt.Println("Hello World - 1 !")

	// 使用自身方法，方法名首字母大小写都可以
	world()
	Say()
}

// 此方法只允许自身调用
func world() {
	fmt.Println("Hello World - 2!")
}

// Say 此方法首字母大写，可以自身调用，也可以外部调用
func Say() {
	fmt.Println("Hello World - 3!")
}
