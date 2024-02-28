package main

import "fmt"

/*
切片追加，make初始化均为0
*/
func main() {
	s := make([]int, 10)

	s = append(s, 1, 2, 3)

	fmt.Println(s)
}
