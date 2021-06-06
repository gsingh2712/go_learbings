package main

import "fmt"

type computer struct {
	brand string
}

/*
How to create functions to create a new pointer
*/
func newcomputer(brand string) *computer {
	return &computer{brand: brand}
}

func main() {

	var com_p *computer // By default it is nil

	if com_p == nil {
		fmt.Println("com_p is nill")
	}

	apple := &computer{brand: "apple"} // How to initialise pointer

	newapplepointer := apple // how to assign

	// How to compare pointers, they compare addresses
	if apple == newapplepointer {
		fmt.Printf("Address are same apple: %p  newapplepointer: %p \n", apple, newapplepointer)
	}

	sony := &computer{brand: "sony"}

	if apple != sony {
		fmt.Printf("Addresses do not compare apple: %p sony: %p \n", apple, sony)
	}

	hp := newcomputer("HP")
	fmt.Printf("Addres for Hp Brand %p  and its value %s \n", hp, *hp)
}
