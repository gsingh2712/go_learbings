package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Buffered io can is used to read from any Input Stream (file , commandline , network)
Here we will see how to read line by line from Standard input

Run
go run main.go
--> Type lines
--> Here
--> ...

Run as below works since Files content are being redirected to std in
go run main.go < File.txt
*/

func main() {

	in := bufio.NewScanner(os.Stdin)
	var lines int

	// in.Scan() function buffers the most line from stand input (in this case), to its own buffer
	// in.Text() function is used to retreive the most recent line scanned
	// in.Err() allows to you to check for error while input Stream scanning

	for in.Scan() { // hover over Scan() and read its doc
		lines++
		fmt.Println("Scanned line ", in.Text(), " Scanned Bytes  ", in.Bytes())
	}

	if err := in.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}
