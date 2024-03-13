package main

import "fmt"

// 函数返回值的初始化
func DeferFunc(i int) (t int) {
	fmt.Println("t = ", t)
	return 2
}

func main() {
	t := DeferFunc(10)
	fmt.Println("t = ", t)
}
