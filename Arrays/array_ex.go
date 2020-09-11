package main

import (
	"fmt"
)

/*
-> Arrays various types of decleration and assignments
-> They are of fixed size and all same type of items
-> Arrays of Arrays
-> Array to Array asssignment,  unlike in other languages, it create two different copies
		-> A moving around array creates independent copies
*/

func main() {
	grades := [3]int{97, 85, 93}     // looks wired if values are already provided at assignment , what is the need for size
	otherway := [...]int{97, 86, 96} // you may avoid to give size as well
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", otherway)
	var students [3]string
	fmt.Printf(" Students : %v  \n", students) // prints empty value
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf(" Students : %v    , Number of Students: %v  \n ", students, len(students))

	/*
		Array of Arrays
	*/
	var idMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1}}
	fmt.Println(idMatrix)

	/*
		Array to Array assigment
		involving creation of a new underlying array
	*/

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	/*
		Array to Array assigment
		here it points to same underlying array
	*/

	c := [...]int{1, 2, 3}
	d := &c
	b[1] = 5
	fmt.Println(c)
	fmt.Println(d)

}
