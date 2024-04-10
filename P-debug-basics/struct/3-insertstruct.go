package main

import "fmt"

type Wheel struct {
	Shape string
}

type Car struct {
	Wheel
	Name string
}

func main() {
	car := &Car{
		Wheel{
			"圆形的",
		},
		"Tesla",
	}
	fmt.Println(car.Name, car.Shape, " ")
}
