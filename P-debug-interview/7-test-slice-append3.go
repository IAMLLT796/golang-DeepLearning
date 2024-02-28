package main

import "fmt"

/*
new和make的区别：
1.都是内存的分配（堆上）
2.make只用于slice、map以及channel的初始化; new用于类型的内存分配，并且内存对应的值为类型的零值
*/
func main() {
	list0 := new([]int)
	*list0 = append(*list0, 1) // list0类型为指针

	list := make([]int, 0)
	list = append(list, 1) // list类型为切片

	fmt.Println(list)
	fmt.Println(list0)
}
