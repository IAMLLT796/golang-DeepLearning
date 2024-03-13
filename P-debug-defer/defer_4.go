package main

import "fmt"

func returnButDefer() (t int) {
	defer func() {
		t = t * 10
	}()
	return 1
}

// 有名函数返回值遇到 defer 的情况
func main() {
	t := returnButDefer()
	fmt.Println("t = ", t)
}
