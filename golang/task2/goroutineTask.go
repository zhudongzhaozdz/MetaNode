package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	var wg sync.WaitGroup
	tasks := []func(){
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务3完成")
		},
	}

	for i, task := range tasks {
		wg.Add(1)
		go func(id int, f func()) {
			defer wg.Done()
			start := time.Now()
			f()
			fmt.Println("任务%d耗时 %v", id+1, time.Since(start))
		}(i, task)
	}
	wg.Wait()
}
