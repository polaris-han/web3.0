package lock_mechanism

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值

func AtomicCounterIncrement() {
	var count int64

	var wg sync.WaitGroup

	// 启动10个协程
	numGoroutines := 10
	wg.Add(numGoroutines)

	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()
	// finalValue := atomic.LoadInt64(&count)
	fmt.Println("最终结果_2:", count)
}
