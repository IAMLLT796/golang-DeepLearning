package main

import "fmt"

func main() {
	classMates1 := make(map[int]string)

	// 添加映射关系
	classMates1[0] = "小明"
	classMates1[1] = "小红"
	classMates1[2] = "小张"

	// 根据 key 获取 value
	fmt.Printf("id is %v\n", 1, classMates1[1])

	// 在声明时进行初始化
	classMates2 := map[int]string{
		0: "小明",
		1: "小红",
		2: "小李",
	}
	fmt.Printf("id %v is %v\n", 3, classMates2[3])
}
