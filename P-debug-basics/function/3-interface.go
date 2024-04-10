package main

import "fmt"

type Printer interface {
	Print(interface{})
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现 Printer 的Print 方法
func (funcCaller FuncCaller) Print(p interface{}) {
	// 调用 funCaller 函数本体
	funcCaller(p)
}

func main() {
	var printer Printer
	// 将匿名函数强转为 FuncCaller 赋值给 printer
	printer = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	printer.Print("Golang is good")
}
