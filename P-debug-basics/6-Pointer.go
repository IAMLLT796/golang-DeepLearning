package main

import (
	"flag"
	"fmt"
)

/*
使用 flag 从命令行中读取参数
*/

func main() {
	// 名称	默认值	提示
	surname := flag.String("surname", "Liu", "您的姓")

	var personName string
	flag.StringVar(&personName, "personName", "LL", "您的名")

	id := flag.Int("id", 0, "您的ID")

	// 解析命令行参数
	flag.Parse()
	fmt.Printf("I am %v %v, and my id is %v\n", *surname, personName, *id)
}
