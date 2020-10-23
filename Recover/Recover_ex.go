package main

import (
	"fmt"
	"log"
)

/*
  Sample program showing the use to defer function for handling recovery caused
  by Panic

  Panic stops the the execution in the flow only in the function where it happens

  Recover funstion is used to recover from panics
  Only usefull in deferred functions

*/

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end") // Executed despite Panic in panciker ()
}

func panicker() {

	fmt.Println(" about to Panic ")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}() // anonymous fucntion no name and are just called
	panic("something bad happened")  // Programmitically called panic button
	fmt.Println(" Done Panciking  ") // not executed since panic happened in this function
}
