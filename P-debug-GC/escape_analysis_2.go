package main

// 当变量的作用域没有超出函数范围时, 就可以放在栈上, 反之就放在堆上

func foo_ea2(arg_val int) *int {
	var foo_val1 *int = new(int)
	var foo_val2 *int = new(int)
	var foo_val3 *int = new(int)
	var foo_val4 *int = new(int)
	var foo_val5 *int = new(int)

	// 此处循环是为了防止 go 编译器将 foo 优化成 inline (内联函数)
	// 如果是内联函数，则 main 调用 foo_ea 将是原地展开，所以 foo_val1-5 相当于 main 作用域的变量
	// 即使 foo_val3 发生逃逸，地址与其他也是连续的
	for i := 0; i < 5; i++ {
		println(&arg_val, &foo_val1, &foo_val2, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
	}

	// 将 foo_val3 返给 main函数
	return foo_val3
}

func main() {
	main_val := foo_ea2(666)
	println(*main_val, main_val)
}
