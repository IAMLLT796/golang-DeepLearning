package main

import "fmt"

// defer 和 return 执行顺序: return 先执行，defer 后执行
func main() {
	returnAndDefer()
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
