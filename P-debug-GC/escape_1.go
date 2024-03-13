package main

func main() {
	data := []interface{}{100, 200}
	println(&data)
	data[0] = 10

	println(&data)
}
