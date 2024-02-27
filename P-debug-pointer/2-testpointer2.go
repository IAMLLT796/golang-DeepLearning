package main

import "fmt"

func main() {
	a := 1
	modify(a)
	fmt.Println(a)

	modifyWithPointer(&a)
	fmt.Println(a)
}

func modify(a int) {
	a = 10
}

func modifyWithPointer(a *int) {
	*a = 100
}
