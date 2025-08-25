package channel

import (
	"fmt"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来

// var channel chan int

func sendData(channel chan<- int) {
	for i := 0; i <= 10; i++ {
		fmt.Println("发送数据：", i)
		channel <- i
	}
	close(channel)
}

func receiveData(channel <-chan int) {
	for v := range channel {
		fmt.Println("接收到数据：", v)
	}
}

func TestChannel() {
	var channel chan int = make(chan int)

	go sendData(channel)

	go receiveData(channel)

	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-channel:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
