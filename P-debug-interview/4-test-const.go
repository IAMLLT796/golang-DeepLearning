package main

import "fmt"

// 常量通常被编译器在预处理阶段直接展开，作为指令数据使用
const c1 = 3

// 变量在运行期分配内存
var d1 = 5

func main() {
	//mt.Println(&c1, c1)
	fmt.Println(&d1, d1)
}
