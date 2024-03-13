package main

import "fmt"

func main() {
	defer_call2()
	fmt.Println("main 正常结束")
}

func defer_call2() {
	defer func() {
		fmt.Println("defer: panic 之前 1, 捕获异常")
		if err := recover(); err != nil {
			fmt.Println(err) // recover() 用于从 panic 中恢复, 组织 panic 继续传播，并获取到 panic 的值
		}
	}()

	defer func() {
		fmt.Println("defer: panic 之前 2, 不捕获")
	}()

	panic("异常内容") // 触发 defer 出栈

	defer func() {
		fmt.Println("panic 之后, 永远执行不到")
	}()
}
