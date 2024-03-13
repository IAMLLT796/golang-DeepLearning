package main

import "fmt"

// defer 遇见 panic, 但是并不捕获异常的情况

func main() {
	defer_call()
	fmt.Println("main 正常结束")
}

func defer_call() {
	defer func() {
		fmt.Println("defer: panic 之前 1")
	}()
	defer func() {
		fmt.Println("defer: panic 之前 2")
	}()

	panic("异常内容") // 触发 defer 出栈

	defer func() {
		fmt.Println("defer: panic 之后, 永远都执行不到")
	}()
}
