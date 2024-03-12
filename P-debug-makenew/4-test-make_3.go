package main

import "fmt"

type Student2 struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*Student2)

	stus := []Student2{
		{Name: "zhou", Age: 19},
		{Name: "wang", Age: 17},
		{Name: "kai", Age: 43},
	}

	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	// 在 foreach 中，v 是结构体的一个副本
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
