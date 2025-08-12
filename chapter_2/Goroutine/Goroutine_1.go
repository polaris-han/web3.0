package goroutine

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数

func print_1(wg *sync.WaitGroup) {
	defer wg.Done()
	var s []int = []int{}
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			s = append(s, i)
		}
	}
	fmt.Println(s)
}

func print_2(wg *sync.WaitGroup) {
	defer wg.Done()
	var s []int = []int{}
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			s = append(s, i)
		}
	}
	fmt.Println(s)
}

func Goroutine_test_1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go print_1(&wg)
	go print_2(&wg)

	wg.Wait()
}
