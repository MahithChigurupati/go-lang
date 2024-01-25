package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("Hello from go routine 1")
		fmt.Println(<-ch)
		wg.Done()
	}(mych, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Hello from go routine 2")
		ch <- 42
		close(ch)
		wg.Done()
	}(mych, wg)

	wg.Wait()

}
