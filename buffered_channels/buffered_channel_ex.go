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
		for i := range ch { // If close(ch) not called as in sender,
			fmt.Println(i) // for range will hit deadlock as it doesn't know where to end
		}
		/*
			for range in above in case of chanel or maps go Reads it like this
			for {
				if i, ok := <- ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}

		*/
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // IMP to close the channel
		wg.Done()
	}(ch)

	wg.Wait()
}
