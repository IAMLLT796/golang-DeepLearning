package main

import "fmt"

func getName() (string, string) {
	return "Liu", "LT"
}

func main() {
	surname, _ := getName()    // 使用匿名变量
	_, personName := getName() // 使用匿名变量
	fmt.Printf("My surname is %v and my personal name is %v", surname, personName)
}
