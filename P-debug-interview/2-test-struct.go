package main

import (
	"fmt"
	"reflect"
)

/*
结构体比较：相同类型的结构体才能进行比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关
*/
func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 22, name: "小刘"}

	sn2 := struct {
		age  int
		name string
	}{age: 22, name: "小刘"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	} else {
		fmt.Println("sn1 != sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 33, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// 结构体中包含map是不可比较的，只有当map都为nil才可以比较
	if sm1.age == sm2.age && reflect.DeepEqual(sm1.m, sm2.m) {
		fmt.Println("sm1 == sm2")
	} else {
		fmt.Println("sm1 != sm2")
	}
}
