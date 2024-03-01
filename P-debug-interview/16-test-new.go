package main

// new用于类型的内存分配，并且内存对应的值为类型的零值
func main() {
	var i *int
	i = new(int)
	*i = 10
	println(*i)
}
