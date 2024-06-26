package main

import "fmt"

func proc(input string, processor func(str string)) {
	// 调用匿名函数
	processor(input)
}

func main() {
	proc("xiaoli", func(str string) {
		for _, v := range str {
			fmt.Printf("%c\n", v)
		}
	})
}
