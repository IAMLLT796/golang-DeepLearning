package main

import "fmt"

/*
nil只能做interface, function, pointer, map, slice, channel的空值
不能做string的空值
*/
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return m[id], true
	}
	return "", false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	res, ok := GetValue(intmap, 3)
	fmt.Println(res, ok)
}
