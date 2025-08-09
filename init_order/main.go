package main

import (
	"fmt"
	// _ "github.com/learn/init_order/pkg1"
)

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

func main() {
	str := []string{"hello", "world"}
	for _, s := range str {
		fmt.Println(s)
	}
}
