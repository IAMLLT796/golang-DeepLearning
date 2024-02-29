package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {
	list = make(map[string]*Student)

	student := Student{Name: "LLT"}

	list["LLT"] = &student

	list["LLT"].Name = "LLT2"

	fmt.Println(list["LLT"])
}
