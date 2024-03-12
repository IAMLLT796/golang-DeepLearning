package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.WaitGroup 是 Go 语言中非常有用的工具，它可以用来等待一组 goroutine 完成

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数退出时调用 Done 来通知 WaitGroup 当前工作已经完成
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // 在启动一个 goroutine 之前, 调用 Add 来增加 WaitGroup 的计数
		go worker(i, &wg) // 启动一个 goroutine
	}
	wg.Wait() // 等待所有工作都完成
	fmt.Println("All workers done")
}
