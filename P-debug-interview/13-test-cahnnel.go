package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			} else {
				fmt.Println("a: ", a)
			}
		}
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
