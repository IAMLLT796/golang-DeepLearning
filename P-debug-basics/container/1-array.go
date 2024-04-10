package main

import (
	"fmt"
)

func main() {
	var classMates1 [3]string
	classMates1[0] = "小明"
	classMates1[1] = "小红"
	classMates1[2] = "小李"
	fmt.Println("The No.1 student is " + classMates1[0])

	classMates2 := [...]string{"小明", "小红", "小李"}
	fmt.Println(classMates2)
}
