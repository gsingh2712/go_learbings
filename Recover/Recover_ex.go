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
	fmt.Println("end") // Executed despite Panic in panciker() ,
	// since defer function has recover function in it , which makes it recover from panic
	// so when control comes back to main function it has recovered from panic

}

func panicker() {

	fmt.Println(" about to Panic ")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}() // anonymous fucntion (no name and are just called with help of paranthesis)
	panic("something bad happened")  // Programmitically called panic button
	fmt.Println(" Done Panciking  ") // not executed since panic happened in this function
}
