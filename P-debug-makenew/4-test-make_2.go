package main

import "fmt"

type Student1 struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*Student1)

	stus := []Student1{
		{Name: "zhou", Age: 19},
		{Name: "wang", Age: 17},
		{Name: "kai", Age: 43},
	}

	for _, v := range stus {
		m[v.Name] = &v
	}
	// 在 foreach 中，v 是结构体的一个副本
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
