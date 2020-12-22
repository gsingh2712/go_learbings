package main

import (
	"fmt"
	"sync"
)

/*
Go Routine is for concurrent programming and parallelism

Synchronization (This is for concurrency)
 -> WaitGroups
 -> Mutexes

-> Add "go" before a function call makes it goRoutine
	-> Tells go to spinOff a green thread and run the function in that thead
	Little discussion on threads in GO
		-> Most programmming languages use OS threads
		   -> It means they have a dedicated stack for handling code executed in that thread
		   -> These theads are very large / expensive which requires concepts of threadPooling
		-> In Go we create Abstraction of thread called GO Routine
		-> Inside Go Runtime a scheduler maps Go Routines on OS threads for a peroid
		-> Scheduler takes turn with every CPU thred and assigns Go Routines a certain amount of processing time on those threads
		-> This abstraction makes Go Routine very lightweight
		-> It is very easy to create 1000's of Go Routine in a Single Runtime , unlike in other languages


-> WaitGroups to wait on the Go Routine

*/

var wg = sync.WaitGroup{} // WaitGroup

func main() {
	fmt.Println("Hello  Go Routines")
	go sayHello()

	wg.Add(1) // One Group to wait on

	var msg = "Closure example:- inner scope acces to outer scope variable"
	go func() {
		fmt.Println(msg)
		wg.Done() // Decrements by one the number of waitgroups being waited on
	}()
	// This is not defined behaviour ,
	// Go Routine will not interrupt main func
	// and it may print updated message
	// this you may say is a race condition
	// alternatively you may change the signature
	// and get string as parameter by value
	msg = "Closure example:- CAVEAT!!! msg updated "
	wg.Wait() // Wait on waitGroups till its count is zero

}

func sayHello() {
	fmt.Println(" sayHello()!! ")
}
