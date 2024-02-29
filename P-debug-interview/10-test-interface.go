package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student1 struct {
}

/*
继承与多态的特点：
1.有interface接口，并且有接口定义的方法
2.有子类去重写interface的接口
3.有父类指针指向子类的具体对象
*/
func (stu *Student1) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var people People = &Student1{}
	think := "love"
	fmt.Println(people.Speak(think))
}
