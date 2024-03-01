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

// new 返回的永远是类型的指针，指向分配类型的内存地址
func main() {
	u := new(user)
	u.lock.Lock()
	u.name = "LLT"
	u.lock.Unlock()

	fmt.Println(u)
}
