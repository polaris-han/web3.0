package channel

import (
	"fmt"
	"sync"
)

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印

func sendData2(channel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		channel <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(channel)
}

func receiveData2(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range channel {
		fmt.Println("接收到数据：", v)
	}
}

func TestChannel_2() {
	var wg sync.WaitGroup
	wg.Add(2)

	var channel chan int = make(chan int, 10)

	go sendData2(channel, &wg)

	go receiveData2(channel, &wg)
	wg.Wait()
}
