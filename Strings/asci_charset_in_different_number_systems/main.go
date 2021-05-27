package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	/*
			 ASCI Values are in 0 to 127 ,
			 since ASCI are just integers
			 you can just give range in integers

			 There are UNICODES that require more than 1 bytes
			 -> 161 - 255 (2 bytes) (utf encoded)
			 -> 9984 - 10175 (3 bytes) (utf encoded)
			 -> 128512 - 128591 (4 bytes) (utf encoded)
			 go run main.go 161 255
			 go run main.go 9984 10175
			 go run main.go 128512 128591
		     Try with the above range and compare the hex and encoded values in those cases, these unicodes
			 are stored in rune types

			 UNICODES can take max upto 4 bytes ,

			 String is a variable length in case of utf8 values. In case a string contains utf8 chars it is
			 difficult to tell the length and chars in string and at which index
			 each characters start. Since utf8 chars can take more than 1 byte

	*/
	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	/*
	 In case no range given OR  Atoi conversion gives 0 , 0
	 use range from A to Z
	*/
	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z' //
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))

	/*
		Printing Ascii Values in Character , decimal , hex and encoded (complete)
		[1] refers to n -> parameter in the Printf
	*/
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", n, string(n))
	}
}
