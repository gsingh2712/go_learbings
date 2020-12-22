package main

import (
	"fmt"
	"sync"
)

/*
Channels are used to pass data between Go Routines and avoid race condition
and memory problems that are difficult to debug

Channel basics :- creating / deleting
Restricting data Channels :- Two way communication / one way communications
Buffered Channels :- Desgin Channel to have internal Data Store, store several messages at once

*/

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(" This is channels")
	ch := make(chan int) // Creating a channel, keyword 'chan' and data type 'int'
	// int tells what type of Data flows
	// Channels are strongly typed can only send data of the type declared
	wg.Add(2)

	go func() {
		i := <-ch // Syntax for receiving data from Channel
		fmt.Println(i)
		wg.Done()
	}() // Receiver Go Routine

	go func() {
		// The BELOW LINE blocks the channel until there is a space
		// availbale in the channel for next data
		ch <- 42 // Syntax for Sending Data to Channel
		// You can pass values from variables to channel as well
		// and its Pass by Value (Copy)
		wg.Done()
	}() // Sender Go Routine
	wg.Wait()

}
