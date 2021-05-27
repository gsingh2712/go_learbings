package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		Some features of utf8 library
		some time utf8 chars are called runes in Golang,
		since rune being typeless and can represent byte , int8 , int16 , int32
		and utf8 chars can take mutliple bytes
	*/
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range strings {
		fmt.Printf("%q\n", s)

		// Print the bytes and rune length of the strings
		// len() print the bytes
		// utf8.RuneCountInString() gives the count of characters
		// since some of chars are utf8 chars, count and bytes might not match
		// len == count in case of
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// %x verb:- Print the hexadecimal number system
		fmt.Printf("\tbytes (hex no. sys)   : % x\n", s)

		// Print the runes / utf8 of the strings in hexadecimal
		// Use % x verb
		//IMP!! Remember for range over a string loops over runes/utf8 chars in the string
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// utf8.DecodeRuneInString(s string) gives the 1st rune and its size from start of string given
		// using the size u can find where the runes ends
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// utf8.DecodeLastRuneInString()
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:]) // Get the byte size of 2nd rune ,
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
	}
}
