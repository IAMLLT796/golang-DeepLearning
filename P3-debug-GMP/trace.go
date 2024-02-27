package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace 的编程过程
// 1.创建文件
// 2.启动
// 3.停止

func main() {
	// 1.创建文件 trace 文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 2.启动
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常要调试的业务
	fmt.Println("Hello GMP")

	// 3.停止trace
	trace.Stop()
}
