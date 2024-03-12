package main

import (
	"fmt"
	"sync"
)

type user struct {
	lock sync.Mutex
	name string
	age  int
}

// new 返回的永远都是类型的指针，指向分配内存的内存地址
func main() {
	u := new(user) // 默认给 u 分配到的内存全部置为 0

	u.lock.Lock() // 可以直接使用，因为 lock 为 0，表示开锁状态
	u.name = "张三"
	u.age = 18
	u.lock.Unlock()

	fmt.Println(u)
}
