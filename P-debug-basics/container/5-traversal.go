package main

import "fmt"

func main() {
	// 数组的遍历
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for index, value := range nums {
		fmt.Println(index, value, " ")
	}
	fmt.Println()

	// 切片的遍历
	slis := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range slis {
		fmt.Println(k, v, " ")
	}
	fmt.Println()

	// 字典的遍历
	Map := map[int]string{
		0: "小明",
		1: "小红",
		2: "小李",
	}
	for k, v := range Map {
		fmt.Println(k, v, " ")
	}
}
