package main

import (
	"fmt"
	"os"
)

/*
Task Mask the http:// links given in test as
commandline input and print the output masking the spam url
without using string matching library functions
*/

const (
	link  = "http://"
	nlink = len(link)
	mask  = 'x'
)

/*
go run main.go "This is trap http://some.com (money)"
*/
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide the input text")
		return
	}

	var (
		input = args[0]
		match = false
		buf   = []byte(input)
		size  = len(input)
	)

	for i := 0; i < size; i++ {

		if len(input[i:]) >= nlink && input[i:i+nlink] == link { // matches pattern 'http://'
			match = true
			i += nlink // jump to character after 'http://'
		}

		switch input[i] {
		case ' ', '\t', '\n': // if these sequences match then following sequence can still contain http://......
			match = false
		}

		// match has been found following , so following characters must be part of URL
		if match {
			buf[i] = mask
		}
	}

	fmt.Println(string(buf)) // Convert byte slice to string
}
