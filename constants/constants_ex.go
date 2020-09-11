package main

import (
	"fmt"
)

/*
Learnings

1. like other longuages if you make Constant variable Upper Case it will exported globally
2. Value assigned to a constant must be determined at compile time and not runtime
		-> Can't assging values like flags , input , expression involving runtime evaluation
3. Inner Constants shadow Outer Constants (Precedene thing)
        -> They are immutable but can be shadowed

*/

const a int16 = 27 // shadow example in a constant

func main() {

	const myConst int = 42 // assingnment at decleration is must
	/*
		const MY_CONST int = 42 // declares and adds it in global scope

	*/
	fmt.Printf(" %v,  %T  \n", a, a)
	const a int = 14 // shadow example  in work here overwritten a's type and value
	fmt.Printf(" %v,  %T  \n", myConst, myConst)
	fmt.Printf(" %v,  %T  \n", a, a)

}
