package main

import "fmt"

type Swimming struct {
}

// 游泳特性
func (swim *Swimming) swim() {
	fmt.Println("swimming is my ability")
}

// 飞行特性
type Flying struct {
}

func (fly Flying) fly() {
	fmt.Println("flying is my ability")
}

// 同时具备飞行和游泳的特性
type WildDuck struct {
	Swimming
	Flying
}

// 只具备游泳特性
type DomesticDuck struct {
	Swimming
}

func main() {
	wild := WildDuck{}
	wild.fly()
	wild.swim()

	domestic := DomesticDuck{}
	domestic.swim()
}
