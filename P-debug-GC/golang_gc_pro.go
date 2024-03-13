package main

func main() {
	main_val := foo(66)
	println(*main_val)
}

func foo(arg_val int) *int {
	var foo_val int = 11

	return &foo_val
}
