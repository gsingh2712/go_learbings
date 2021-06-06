package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type permissions map[string]bool // #3
// All member variables are exportable (First letter capital) to allow Marshalling / Encoding
type user struct { // #1
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"` // #6

	// name        string // #1
	// password    string // #1
	// permissions // #3
}

func main() {
	users := []user{ // #2
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "66", permissions{"write": true}},
	}

	out, err := json.Marshal(users) // #4 // returns byte slice
	if err != nil {                 // #4
		fmt.Println(err)
		return
	}
	fmt.Println("Printing simple Marshalled Json")
	fmt.Println(string(out)) // Converting byte slice to string

	out, err = json.MarshalIndent(users, "*", " -> ") // #5
	if err != nil {                                   // #4
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Repeat("-", 45))
	fmt.Println("Printing Different Indented Marshalled Json uisng * as Prefix and ' -> ' as indent")
	fmt.Println("Notice Field Permission is dropped in json (omitempty)")
	fmt.Println(string(out)) // #4
	fmt.Println(strings.Repeat("-", 45))
	fmt.Println("Printing Indented Marshalled Json")
	fmt.Println("Notice Field Permission is dropped in json (omitempty)")
	out, err = json.MarshalIndent(users, " ", "\t") // #5
	if err != nil {                                 // #4
		fmt.Println(err)
		return
	}
	fmt.Println(string(out)) // #4
}
