package main

import "fmt"

func main() {
	// 引用类型的变量，不仅要声明他，还要为他分配内存空间
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)
}
