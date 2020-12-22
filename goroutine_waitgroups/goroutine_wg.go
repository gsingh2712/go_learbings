package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

// wg Waitgroup is being used/ shared via closure feature and exposed to other functions
// Waitgroup is safe and meant to be used concurrenly across multiple Go Routines

/*
 The Result will be undefined since all Go Routines are executing independently
 and schdeuled independendently leading to Race Condition

 This can be solved by Mutexes
*/

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) // Total 20 Wg's
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
