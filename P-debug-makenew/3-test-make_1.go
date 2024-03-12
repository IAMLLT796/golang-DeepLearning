package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {
	list = make(map[string]Student)

	student := Student{"LLT"}
	list["student"] = student
	// 不能像下面这样做，因为是一个值传递过程
	// list["student"].Name = "LLT0"
	tmp := list["student"]
	tmp.Name = "LLT0"
	list["student"] = tmp

	fmt.Println(list["student"])

	// list["student"] = &student
	// list["student"].Name = "LLT1"

	fmt.Println(list["student"])
}
