package main

import (
	"fmt"
	"sync"
)

/*

Example where both Go routines act as Reader Wirter
GO_2 writes 42 in ch and is read and printed by GO_1

So both Go routines act as reader and writer

To avoid that we need to add direction bias in channels as parameters

*/

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	wg.Add(4)
	go func() { // GO _1
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() { // GO_2
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	go func(ch2 <-chan int) { // GO_3 Receive Only Channel
		i := <-ch2
		fmt.Println(i)
		// Adding ch_2 <- 100 is syntax error since we are trying to send to channel
		wg.Done()
	}(ch2)
	go func(ch2 chan<- int) { // GO_4 Send Only Channel
		ch2 <- 50
		wg.Done()
	}(ch2)
	wg.Wait()
}
