package main

import "fmt"

func main() {
	str := "hey"
	bytes := []byte{104, 101, 121}

	/*
		A rune is typeless values and its default type is rune
		and depending on context it can take 1byte , 2byte , 3byte , 4byte

		Hence in case of string with utf8 chars rune is prerefered to read each char.
	*/

	// same as: []byte("hey")
	fmt.Printf(`"hey" as bytes   : %d`+"\n", []byte(str))

	// same as: string([]byte{104, 101, 121})
	// in golang you can  convert bytes into strings using string function
	fmt.Printf("bytes as string  : %q\n", string(bytes))

	// runes are unicode codepoints (numbers)
	fmt.Println()
	fmt.Printf("%c                : %[1]d\n", 'h')
	fmt.Printf("%c                : %[1]d\n", 'e')
	fmt.Printf("%c                : %[1]d\n", 'y')

	// a rune literal is typeless
	// you can put it in any numeric type
	var (
		anInt   int   = 'h'
		anInt8  int8  = 'h'
		anInt16 int16 = 'h'
		anInt32 int32 = 'h'

		// rune literal's default type is: rune
		// so, you don't need to specify it.
		// aRune   rune  = 'h'
		aRune = 'h'

		// and so on...
	)

	fmt.Println()
	fmt.Printf("rune literals are typeless:\n\t%T %T %T %T %T\n",
		anInt, anInt8, anInt16, anInt32, aRune)

	fmt.Println()

	// all are the same rune

	// beginning with go 1.13 you can type: 0b0110_1000 instead
	// fmt.Printf("%q as binary: %08[1]b\n", 0b0110_1000)
	/* [1] --> parameter position in Printf() which below refers to 104, 'h' and x068 in
	below Printf statements

	PRINTING 'h' in different number systems
	*/
	fmt.Printf("%q in decimal: %[1]d\n", 'h')   // %d -> tells to print decimel format
	fmt.Printf("%q in binary : %08[1]b\n", 'h') // %b -> tells to print in binary 8 digit format 08 asks to be paaded with zero leading digits

	fmt.Printf("%q in hex    : 0x%[1]x\n", 'h') // %x  -> tells to print in hex format
}
