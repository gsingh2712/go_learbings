package main

import (
	"fmt"
)

/*
Iota is a special of enumerated constants , it is like a counter and automatically increcments
can be helpful in creating an enumeration
*/

const (
	a = iota
	b = iota
	c = iota
)

const (
	e = iota
	f
	g
)

/*
 interesting use of iota , in bit shifting
*/

/*
	ignore the first value by assigning to blank identifier
	_ is write only garbage store in go for storing default value
*/

/*
   Enumerated constants along with iota and bit shiting
   can work real wonders
*/

/*
   It uses the same function to assign further values
   as used in the initial function
   1 << ( 10 * iota) is the function in the case below
*/

const (
	// _ is write only garbage store in go for storing default value in this case 0
	_ = iota
	// KB kilobyte
	KB = 1 << (10 * iota) // 2 ^10
	// MB megabyte
	MB // 2 ^ 20
	// GB gigabyte
	GB // 2 ^30
	// TB Terabyte
	TB // 2 ^ 40
)

func main() {

	/*
	  here it shows that value of iota auto incremented in first block
	*/

	fmt.Printf(" %v, %T\n", a, a)
	fmt.Printf(" %v, %T\n", b, b)
	fmt.Printf(" %v, %T\n", c, c)

	/*
	  here we see value of varibales auto incremented
	  and iota is scoped to constant block
	  it is reset back to 0 always in a new block
	*/

	fmt.Printf(" %v, %T\n", e, e)
	fmt.Printf(" %v, %T\n", f, f)
	fmt.Printf(" %v, %T\n", g, g)

	fmt.Printf(" KB %v bytes, %T\n", KB, KB)
	fmt.Printf(" MB %v KB , %v bytes, %T\n", MB/KB, MB, MB)
	fmt.Printf(" GB %v MB,  %v bytes,   %T\n", GB/MB, GB, GB)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB \n", fileSize/GB)

}
