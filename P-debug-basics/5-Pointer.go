package main

import "fmt"

func main() {
	// 声明一个 string 类型
	str := "Golang is Good!"

	// 获取 str 的指针
	strPtr := &str

	fmt.Printf("str type is %T, and value id %v\n", str, str)
	fmt.Printf("strPtr type is %T, and value id %v\n", strPtr, strPtr)

	newStr := *strPtr // 值拷贝
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr, newStr, &newStr)

	*strPtr = "Java is good too" // 通过指针对变量进行赋值
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr, newStr, &newStr)
	fmt.Printf("str type is %T, value is %v, and address is %p\n", str, str, &str)
}
