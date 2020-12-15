package main

import (
	"fmt"
)

/*
-> Interfaces are like type
-> Interfaces define behaviour
		-> We store method definitions in Interfaces
-> Interfaces allow you to have Polymorphic behaviour
	-> In Below Example you can have TCPwriter, FireWriter ...
		as writer var and call Write() functions and corresponding implementation will be called
-> In Go we implicitly implement interface
	-> Not sure what it means , seems not something to come acorss often

-> Smaller interfaces / Interface with single method is best practise

-> Within one interface you can club multiple other interface

type Closer interface {
	Close() error
}

type Writer interface {
	Write() error
}

type WriterCloser interface {
	Writer
	Closer
}

*/

func main() {
	fmt.Println("Interface Demo")
	var w writer = consoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
