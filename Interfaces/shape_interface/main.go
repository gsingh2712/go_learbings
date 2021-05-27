package main

import (
	"fmt"
	"math"
)

/*
Interface in GO are a Type!!!!
Every Shape must implement area() interface
*/
type shape interface {
	area() float64
}

// another shape square
type square struct {
	side float64
}

// another shape circle
type circle struct {
	radius float64
}

/*
Square implements area function
*/
func (s square) area() float64 {
	return s.side * s.side
}

/*
Circle implements area function
*/
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
Info accepts only Shape
*/
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}

	/*
		square and circle are shapes as they implemented area() method and
		hence can be sent to info () functions
	*/
	info(s)
	info(c)
}
