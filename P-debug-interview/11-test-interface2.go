package main

import "fmt"

// 非空接口: 具有tab和data两个字段，tab指向具体类型的方法表，data指向具体类型的数据
type People1 interface {
	Show()
}

type Student2 struct {
}

func (stu *Student2) Show() {

}

func live() People1 {
	var stu *Student2
	return stu // stu是一个指向nil的空指针，return时会触发People = stu的隐式转换，所以返回的是一个空的People接口
}

// 结果：2222222
func main() {
	if live() == nil {
		fmt.Println("1111111")
	} else {
		fmt.Println("2222222")
	}
}
