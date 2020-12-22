package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

Continued from Waitgroup Example how to solve race condtion via mutex

*/

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Read Write Mutex,
// as many as can read the data but only one can write at a time
// and if anything is reading we cannot write
// ==> infinite number of Readers and onle one writer
// -->  When Writer comes to write it will wait for readers to complete
// and take a write lock

// Using mutexes to synchronize things in this example

// YOU CAN CHECK for race conditions in your Go Application
// using go compiler with -race param to (go run / build / install  ) whatever you use to run Go Application
// example  go run -race <GoCode>

func main() {
	/*
		IT is for getting number for threads ,
		by default GO gives #threads as #cores (to get that number print  ==> runtime.GOMAXPROCS(-1))

		GOMAXPROCS serves as tuning parameter taking too high a number can adversely impact
		as you might have studied in OS concepts (too much expensive)(more memory , schedudler overburdened)


	*/
	runtime.GOMAXPROCS(100)

	/*
		Locks have to outside Go Routines
		to synchronize across print and increment
		if instead they are put inside  though increments will be in sync
		but prints will be unreliable

		Reason for it is since each go routine is independent and scheduled
		independently subseqeuent sayHello are getting sheduled first while
		increments keep waiting, so take a write Lock outside increment so that
		subsequent Read locks have to wait.

		Pay attention to the blob and see how it achieves sync

		Food for thought?
		Though this solves problem but it has destryed concurrecny and
		parallelism , it would have been better to run without Go Routines at all.
		This is just an example to illustrate how to use mutexes
	*/
	for i := 0; i < 10; i++ {
		wg.Add(2) // Total 20 Wg's
		m.RLock() // read lock
		go sayHello()
		m.Lock() // write lock
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
