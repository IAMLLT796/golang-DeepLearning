package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// 两个切片在append的时候，需要对第二个切片进行...打散再拼接
	s1 = append(s1, s2...)

	fmt.Println(s1)
}
