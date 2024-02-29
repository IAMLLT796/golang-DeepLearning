package main

type S struct {
}

func f(x interface{}) {

}

// 只能接受 *interface{}
func g(x *interface{}) {

}

// Golang是强类型语言, interface是所有golang类型的父类, 所以可以接受任何类型的参数，包括指针类型
func main() {
	s := S{}
	p := &s
	f(s)
	//g(s)
	f(p)
	//g(p)
}
