package lock_mechanism

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	// c.mu.Unlock()
}

func TestCounter() {
	var c Counter
	var wg sync.WaitGroup

	// 启动10个协程
	numGoroutines := 10
	wg.Add(numGoroutines)

	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range 1000 {
				c.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终结果：", c.count)
}
