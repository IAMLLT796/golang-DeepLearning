package main

import (
	"fmt"
	"unicode/utf8"
)

func traverse(text string) {
	// 会出现乱码
	for _, g := range []byte(text) {
		fmt.Printf("%c", g)
	}

	// 按照字符遍历字符串
	for _, h := range text {
		fmt.Printf("%c", h)
	}
}
func main() {
	f := "Golang编程"
	fmt.Printf("byte len of f is %v\n", len(f))
	fmt.Printf("rune len of f is %v\n", utf8.RuneCountInString(f))

	traverse(f)
}
