package main

import "fmt"

/*

Struct is type of Collection type,
Since in itself it can store different types of data types
While all other Collection types store data of same types

-> Inspecting elments you need to use . dots

-> Structs work by passing copies and not by reference, unless you use 'address &' of operator
	-> Structs are value types

-> Go doesn't support inheritence , but supports embedding of structs to achieve this
	-> To achieve 'HAS A' relationship behaviour of inheritance and not 'IS A' relationship
	-> But there types are still independent, not that struct being embedded becomes base class
	-> Embedding has very limited use in terms of Polymorphism and should be avoided
	-> Generally Interfaces should be used to describe common behaviour

Naming conventions of Structs follow same things as that of variables
1. UpperCase in start means it is visible outside the scope of package
	* Filed names have to be UpperCase 1st letters so it can be accessed outside package
2. LoweCase in start it means limited to the package
3. Anonymous structs where you dont need to declare in the form of typedef ... struct {...}
	-> It can't be exported out of package

*/

type doctor struct {
	number     int
	actorName  string
	companions []string
}

// Doctorgl global decleration
type Doctorgl struct {
	/*
	  Making member variables uppercase Global access possible
	*/
	Number     int
	ActorName  string
	Companions []string
}

// Animal Global
type Animal struct {
	Name   string
	Origin string
}

/*
embeding example on how to achieve inheritance
*/

// Bird Global
type Bird struct {
	Animal
	SpeedKPH float32
	Canfly   bool
}

func main() {
	aDoctor := doctor{
		number:    4,
		actorName: "SSR",
		companions: []string{ // note it is a slice of string
			"Liz",
			"Jo",
			"Sarah",
			"Smith",
		},
	}

	fmt.Println(aDoctor)
	// inspecting elements
	fmt.Println(aDoctor.companions[1:]) // using slice feature to print from index 1 to last

	fmt.Printf("Structs with member and its values  %+v  \n", aDoctor)

	glDoctor := Doctorgl{
		Number:    4,
		ActorName: "SSR",
		Companions: []string{ // note it is a slice of string
			"GL Liz",
			"GL Jo",
			"GL Sarah",
			"GL Smith",
		},
	}

	fmt.Println(glDoctor.Companions[2:])

	/*
		Anonymous Struct Decleration example with an initializer syntax
		No need to create like typedef <> struct {...}
	*/
	anonStructexample := struct{ name string }{name: "Anonymous Struct Example"}
	fmt.Println(" anonStruct  ", anonStructexample)

	/*
			 Embedding example to achieve inheritance
			 Embedding Animal in bird does'nt mean that Bird's struct type has become that of Animal
		     Unlike in java , c++ Bird and Animal are still independent types
	*/
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.Canfly = false
	fmt.Println("  Bird emu  ", b)

	c := Bird{}
	c.Name = "Emu"
	c.Origin = "Australia"
	c.SpeedKPH = 48
	c.Canfly = false
	fmt.Println(" Bird b and c are same  ", b == c)
	/*
	 structs can be compared if they dont contain uncomparable fields like slice , map , function ..
	*/
}
