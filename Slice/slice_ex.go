package main

import (
	"fmt"
)

/*
  -> Slice are naturally called reference type,
	 -> Caution needed in case you dont want undelrying reference updated randomly
	 -> Slices work on reference
  -> More or less looks like Python slices operationally
  -> Another way of creating slices is via make command
  -> Slice unlike array are not limited by size and you keep appending vaalues
			-> creats new array to accomodate for more space
			-> new capcicty is always in power of 2 no matter how much length
			-> 0 , 1, 2 , 4 , 8


*/

func main() {
	a := []int{1, 2, 3} // this is a slice , note how change in assignment makes it slice and  not arrays
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Length: %v \n", cap(a))

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // this is a slice , note how change in assignment makes it slice
	d := c[:]
	e := c[3:]  // slice from 4th till end
	f := c[3:6] // slice 4 , 5 , 6 th element
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	g := make([]int, 3, 100) // slice of 3 or len but capacity 100
	fmt.Println(g)
	fmt.Printf("Length: g %v\n", len(g))
	fmt.Printf("Capacity: g %v\n", cap(g))

	h := []int{}
	fmt.Println(a)
	fmt.Printf("Length: h %v \n", len(h))
	fmt.Printf("Capacity: h %v \n", cap(h))

	h = append(h, 1)
	fmt.Printf("Length: h %v \n", len(h))
	fmt.Printf("Capacity: h %v \n", cap(h))

	h = append(h, 2, 3)
	fmt.Printf("Length: h %v \n", len(h))
	fmt.Printf("Capacity: h %v \n", cap(h))

}
