package main

import (
	"fmt"
)

/*
It is a contol flow construct
Defer :-
    Invoke a function call but delay its execution to some future point in time
Panic :-
   Application running into a state where it can't run / recover from , this scenario can be triggred
   both by application and programmitically

   Panic are executed after defer are executed

  Panic has different connotation than Exception , since many situations that cause Exception
  in other languages are actually normal situation in GO and it instead for those situations
  just throws Error and not Exception , For Ex:- File Not Found is exception in Java

Recover :-
   Saving the program when it triggers Panic

*/

/*
  Defer works by executing the function passed to it, just before return/exit
  of the function inside which Defer has ben called

  Defer functions are executed as in LIFO order

  Used to group open and close functions together
			-> Be Careful in loops

  Defers run in LIFO Order

*/

func main() {

	fmt.Println("Start")
	defer fmt.Println("middle") // will be executed just before exiting main
	fmt.Println("end")

	/*
	  Notice the LIFO order of executed functions
	*/
	defer fmt.Println("Start_1")
	defer fmt.Println("middle_1")
	defer fmt.Println("end_1")

	/*
	  Defer Riddle
	  What valule of test is printed
	*/
	test := "start"
	defer fmt.Println("Prints the value of test at the time of function call and not the execution  ", test)
	test = "end"

	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

}
