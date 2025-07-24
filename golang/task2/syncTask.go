package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取计数
func (c *SafeCounter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

	// counter := SafeCounter{}
	// var wg sync.WaitGroup
	// // 启动100个goroutine同时增加计数
	// for i := 0; i < 10; i++ {

	// 	wg.Add(1)
	// 	go func(id int) {
	// 		//通知等待组当前协程已完成
	// 		defer wg.Done()
	// 		fmt.Printf("协程 %d 开始\n", id)
	// 		for j := 0; j < 1000; j++ {
	// 			counter.increment()
	// 		}
	// 		fmt.Printf("协程 %d 完成\n", id)
	// 	}(i)
	// }
	// wg.Wait()
	// time.Sleep(time.Second)

	// fmt.Println("Final count: %d", counter.getCount())

	// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var initial int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("协程 %d 启动", id)
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&initial, 1)
			}
			fmt.Println("协程 %d 递增完成", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("Final count: %d", initial)
}
