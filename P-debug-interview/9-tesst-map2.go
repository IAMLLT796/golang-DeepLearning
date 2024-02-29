package main

import "fmt"

type student struct {
	Name string
	Age  int64
}

func main() {
	m := make(map[string]*student)

	// 定义student数组
	stus := []student{
		{Name: "小赵", Age: 19},
		{Name: "小望", Age: 24},
		{Name: "小刘", Age: 22},
	}

	// 将数组依次添加到map中
	// 问题：stu是一个临时变量，每次循环都会被重新赋值，所以最终map中的值都是最后一个stu的值
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}

	//for i := range stus {
	//	m[stus[i].Name] = &stus[i]
	//}

	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	// 打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
