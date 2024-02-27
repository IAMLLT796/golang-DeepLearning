package main

/*
函数返回值需要有命名
*/
func myFunc(x, y int) (sum int, err error) {
	return x + y, nil
}
