package main

import (
	"fmt"
	"sync"
)

/*
make 与 new 的异同
1.都是用来堆空间分配内存的
2.make 只用于 slice, map, channel 的初始化，无可替代
3.new 用于类型内存分配， 初始值为0，不常用
*/
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
