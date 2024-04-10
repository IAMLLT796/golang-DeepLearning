package main

import "fmt"

type Cat interface {
	CatchMouse()
}

type Dog interface {
	Bark()
}

type CatDog struct {
	Name string
}

// 实现 Cat 接口
func (this *CatDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n", this.Name)
}

// 实现 Dog 接口
func (this *CatDog) Bark() {
	fmt.Printf("%v barked loudly!\n", this.Name)
}

func main() {
	var cat Cat

	cat = &CatDog{
		"Tom",
	}
	cat.CatchMouse()

	var dog Dog
	dog = &CatDog{
		"Jerry",
	}
	dog.Bark()
}
