package main

import "fmt"

func main() {
	a := 10 // int类型
	b := &a // *int类型
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, b)
	// 变量b是一个int类型的指针，它存储的是变量a的内存地址
	c := *b // 根据内存地址取值
	fmt.Println(c)
}
