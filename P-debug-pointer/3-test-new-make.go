package main

import "fmt"

/*
new 和 make 的区别
1.二者都是用来做内存分配的
2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
3.new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
*/
func main() {
	var a *int // 没有初始化
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int)
	b["two"] = 1
	fmt.Println(b)
}
