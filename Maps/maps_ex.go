package main

import "fmt"

/*
 1. Map always map one key type to one value type
 2. Key type should be something that can be tested for equality
	  -> slices , maps and functions can't be used as key type
 3. Return Order in a map is not gauranteed
 4. Maps in assignment and passing by are by reference , hence manipulating one changes the underlying map
 5. Comma , Ok syntax to find non exitant keys

*/

func main() {
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
			m := map[[]int]string{}
		    This is an illegal map decleration since key is of type slice
	*/
	fmt.Println(" State Populations ", statePopl)
	/*
	  You can declare map via make statement
	*/
	statepl := make(map[string]int)
	statepl = map[string]int{
		"UP": 20,
	}
	fmt.Println(" State Populations ", statepl)

	/*
	  Notice the Ordering in map
	*/
	// Accessing map values
	fmt.Println(" State Populations ", statePopl["Ohio"])
	// Adding values
	statePopl["Georgia"] = 21
	fmt.Println(" State Populations ", statePopl)
	delete(statePopl, "Georgia")
	fmt.Println(" State Populations ", statePopl)

	// How to check for non existant keys
	pop, ok := statePopl["Oho"] // comma Ok syntax for checking non existant keys
	fmt.Println(pop, ok)        //

	_, ok = statePopl["Oho"] // comma Ok syntax for checking non existant keys
	fmt.Println(ok)          //

	// Count of elements in map with len
	fmt.Println(len(statePopl))

	/*
	  Pass by reference example of maps
	  Ohio deleted in both
	*/

	sp := statePopl
	delete(sp, "Ohio")
	fmt.Println(" Changed map ", sp)
	fmt.Println(" Underlying map ", statePopl)

}
