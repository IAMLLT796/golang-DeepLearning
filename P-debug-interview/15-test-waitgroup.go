package main

import "sync"

var wg1 = &sync.WaitGroup{}

func main() {
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func(i int) {
			println(i)
			defer wg1.Done()
		}(i)
	}
	wg1.Wait()
}
