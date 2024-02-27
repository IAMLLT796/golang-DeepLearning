package main

import (
	"fmt"
	"sync"
	"time"
)

var set = make(map[int]bool, 0)

func printOne(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

var m sync.Mutex

// 使用互斥锁
func printOneWithMutex(num int) {
	m.Lock()
	// defer m.Unlock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}
func main() {
	for i := 0; i < 10; i++ {
		go printOneWithMutex(100)
	}
	time.Sleep(time.Second)
}
