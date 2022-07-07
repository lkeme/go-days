package main

import "fmt"

const pi = 3.1415926

const (
	a    = iota // 0
	b           // 1
	c    = 100  // 100
	d           // 注意这里已经中断 是参考上一个值  是100
	e    = iota // 再次使用iota 继续iota枚举 所以是4
	f, g = iota, iota
)

const (
	j = iota + 100
	k
	l
)

const (
	x = iota + 3
	xx
	xxx
	y = iota + 6
	yy
	yyy
	z = iota + 9
	zz
	zzz
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(j, k, l)
	fmt.Println(x, xx, xxx, y, yy, yyy, z, zz, zzz)
}
