package main

import (
	"fmt"
)

/*
  -> Funnctions decleration need to have keyword ->  func
  -> Name of the function can be in upper / lower case depending upon visibility across packages
  -> Arguments in the function are like just any variable decleration
  -> Return type is defined after parenthesis
  -> Go can only returns pointer to local variables
  -> You can return as many values from Go Functions,  like Error as well as return value

  --> Functions in go are actually very powerful and can be treated as TYPES
			-> It means they can be assigned to a variable
*/

func main() {
	sayMessage("Hell Go|")

	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Function assignment to a variable
	f := func() {
		fmt.Println("Hello Go! Anonymous")
	}
	f() // --> Calling the function via assgigned variable
}

func sayMessage(msg string) {
	fmt.Println(msg)
	s := sum("The Sum is ", 1, 2, 3, 4, 5)
	fmt.Printf(" Returned value %v \n", s)
}

// varitic Parameters
// variatic Paremeters must form the last set of params to the functions

func sum(msg string, values ...int) int {
	// variable procedding with dots tell it compile all the arguements
	// into a slice
	fmt.Println(values) // prints the slice
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
