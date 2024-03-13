package main

func main() {
	ch := make(chan []string)
	s := []string{"LLT"}

	go func() {
		ch <- s
	}()
}
