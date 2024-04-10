package main

import "fmt"

type Person struct {
	Name     string // 姓名
	Birthday string // 生日
	ID       int64  // 身份证号
}

// 指针类型, 修改个人信息
func (person *Person) changeName(name string) {
	person.Name = name
}

// 非指针类型, 打印个人信息
func (person Person) printMess() {
	fmt.Printf("My name is %v, and my birthday is %v, and my id is %v\n", person.Name, person.Birthday, person.ID)

	// 尝试修改个人信息, 但是对原接收器并没有影响
	// person.ID = 3

}

func main() {
	p1 := Person{
		Name:     "CCC",
		Birthday: "1998-11-22",
		ID:       1,
	}
	p1.printMess()

	p1.changeName("DDD")
	p1.printMess()
}
