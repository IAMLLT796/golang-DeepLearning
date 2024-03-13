package main

import "fmt"

func foo_6(a *int) {
	return
}
func main() {
	data := 10
	f := foo_6
	f(&data)
	fmt.Println(f)
}
