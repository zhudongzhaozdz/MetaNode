package main

import (
	"fmt"
	"sync"
)

func main() {
	// printChannel()
	bufferedChannel()
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func printChannel() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			fmt.Println("放入channel中的值: %d", i)
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 从channel接收数据
		for value := range ch {
			fmt.Println("从channel中接收到的值:%d", value)
		}
	}()
	wg.Wait()
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func bufferedChannel() {
	fmt.Println("====带缓冲channel示例====")
	ch := make(chan int, 10) // 缓冲大小为2
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 1; i < 101; i++ {
			fmt.Println("放入channel的值:%d", i)
			ch <- i
		}
	}()
	wg.Add(1)

	go func() {
		// 从channel接收数据
		defer wg.Done()
		for value := range ch {
			fmt.Println("从channel中接受到的值:%d", value)
		}
	}()
	wg.Wait()
}
