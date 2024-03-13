package main

import "fmt"

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}

// 4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// 1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// 3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

// 0
// 2
func DeferFunc4() (t int) {
	// 赋值 defer 中的 func 入参 t 为 0
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}
