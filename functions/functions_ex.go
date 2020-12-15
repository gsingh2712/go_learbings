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

  --> See how method decleration is done for Structs
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

	gInstance := greeter{
		greeting: "Hello From Greeter",
		name:     "Go",
	}

	gInstance.greet()

	gInstance.referencegreet() // Here gInstance is passed by pointer
	fmt.Println(gInstance.greeting, gInstance.name)
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

type greeter struct {
	greeting string
	name     string
}

/*
This is method greet() declared for struct greeter
g represents the context / struct instance on which greet() has been called
The g here is pass by value any editing on Struct variable wont reflect
*/
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.greeting = "" // This wont work since it is pass by value
}

/*
This is another method declared for struct greeter
The g here is pass by pointer any editing on Struct variable will reflect
*/
func (g *greeter) referencegreet() {
	// Go supports implicit dereferencing hence no need for (*g).greeting
	g.greeting = "Hello from Referenced greet" // This works since it is pass by reference
}
