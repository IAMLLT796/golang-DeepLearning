package main

import "fmt"

func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	return func() int {
		initial++
		// 返回当前计数
		return initial
	}
}

func main() {
	// 计数器1
	c1 := createCounter(1)

	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	// 计数器2
	c2 := createCounter(10)
	fmt.Println(c2()) // 11
	fmt.Println(c1()) // 4
}
