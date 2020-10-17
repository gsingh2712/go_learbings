package main

import (
	"fmt"
)

/*

1. Exit from infinite loop using break keyword
	-> break allows exit from nearest loop
2. Notice how to use break from inner/nested loops using LABEL
3. Working with Collections using for range loop

*/

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Notice the syntax for using double variables in for loop
	// Take notice of the initial assignment and the increment operations on i and j
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// Example of Label
Loop: // Label is marked for the outermost Loop i
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop // break Label allows to exit from Loop for which it is marked
				// in ths case it is marked for i
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)

	// Example of using range
	for k, v := range s { // key and value
		fmt.Println(k, v)
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

	for k, v := range statePopl {
		fmt.Println(k, v)
	}

	// Use Looping on String
	st := "Hello Go!"
	for k, v := range st {
		fmt.Println(k, "Ascii value", v, " Character ", string(v)) //
	}

}
