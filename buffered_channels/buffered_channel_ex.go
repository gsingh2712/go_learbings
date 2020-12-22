package main

import (
	"fmt"
	"sync"
)

/*
Buffered Channels are used when you want / might send more data at a time
than can be consumed
Sender and Receiver operate at different speed
*/

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()
}
