package main

import (
	"fmt"
)

/*
  -> Curly braces are compulsory in if statements

*/

func main() {

	if true {
		fmt.Println(" This test is true")
	}

	if false {
		fmt.Println(" This test is false")
	}

	statePopl := map[string]int{ // check how to specify key and value type
		"cali":     12,
		"Texas":    13,
		"Florida":  14,
		"NY":       15,
		"Pensyl":   16,
		"Illinois": 17,
		"Ohio":     18,
	}

	/*
	   How to use comma Ok syntax of maps in if statement
	*/
	if pop, ok := statePopl["Florida"]; ok {
		// pop variable is in just the scope of this if statement
		fmt.Println(pop)
	}

	/*
	   Example for control flow operators
	*/

	number := 50
	guess := 30

	if guess < number { // no enclosing in paranthesis needed
		fmt.Println("Too low")
	}

	if guess > number {
		fmt.Println("Too low")
	}

	// Logical Operators

	if guess >= 1 && guess < 100 {
		fmt.Println("  Guess must be between 1 and 100  ")
	}

	// not operator
	fmt.Println(!true)

	// Swtich Example
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	// Switch Cases with Ranges

	switch 3 {
	case 1, 2, 3:
		fmt.Println(" 1 to 3 ")
	case 4, 5, 6:
		fmt.Println("4 to 6")
	default:
		fmt.Println("Default")
	}
	// Switch Tagless syntax

	i := 10
	/*
	  Despite Overlap between cases in switch
	  only former is executed
	  break key is implied by default in every switch case to change behaviour use fallthrough

	*/
	switch {
	case i <= 10:
		fmt.Println("Less than equal to 10")
	case i <= 20:
		fmt.Println("Less than equal to 20")
	default:
		fmt.Println("Greater than 20")
	}

	// Caveat : fallthrough means coder is responsible for correctness
	// next statement of following  case  is executed by default
	switch {
	case i <= 10:
		fmt.Println("Less than equal to 10")
		fallthrough
	case i <= 20:
		fmt.Println("Less than equal to 20")
	default:
		fmt.Println("Greater than 20")
	}
}
