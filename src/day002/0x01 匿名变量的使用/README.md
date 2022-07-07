## 0x01 匿名变量的使用

上一节我们说到，变量一旦在函数里定义后，必须要使用它，否则报错。  
但是很多情况下，有一些变量在上下文之间使用不到，此时就要用到匿名变量。

```go
package main

import "fmt"

func info(content string) (msg string, err error) {
	// 会返回两个值
	return content + "\r\n", nil
}

func main() {
	// 在不同的业务条件下，可能有某个参数用不上，不需要下文判断或者使用
	msg, _ := info("Hello world")
	// 使用下划线_替代变量名
	fmt.Println(msg)
}
```

> 匿名变量不占用内存空间，不分配内存，可以多次声明使用

## 链接

- [目录](../../../README.md)
- 上一节：[变量与常量的定义](../0x01%20匿名变量的使用)
- 下一节：
