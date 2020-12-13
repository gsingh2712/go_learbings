package main

import "fmt"

/*

-> Go doesn't allow pointer airthmentic by default as in C and C++.
(&arr[1] -4)  is invalid in Go.

-> If you really want to use pointer airthmetic , use package "unsafe"

-> printing pointers behaves differently in case of Struct

-> Every initialised  pointer in Go store value <nil>

*/

type myStruct struct {
	foo int
}

func main() {
	var a int = 42  // declerative varibale decleration
	var b *int = &a // pointer decleration of type int
	// pointers store address
	fmt.Println(&a, b) // address of a is same as value of b
	fmt.Println(a, *b) // here * acts as dereferencing operator

	// Pointer Airthmetic example

	arr := [3]int{1, 2, 3}
	brr := &arr[0]
	crr := &arr[1] // its value will be 8 bytes higher than brr

	fmt.Printf("%v %p  %p\n", arr, brr, crr)

	var ms *myStruct                                            // by default get initialised by nil
	fmt.Println(ms)                                             // print <nil>
	ms = &myStruct{foo: 42}                                     // Or new(myStruct)  new doesn't give option of value initialization
	fmt.Printf(" ms value %p  foo value %v  \n", ms, (*ms).foo) // print like &{42}
	// Printing foo without dereferencing also works, it is compiler optimization added
	fmt.Printf(" Print foo without dereferencing %v  \n", ms.foo)
}
