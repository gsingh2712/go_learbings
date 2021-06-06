package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`

	/*
		Device []struct { // Nested struct and its marshalling and unmarshalling
			Name    string
			Battery int
		} `json:"devices"`
	*/

	Devices []Device `json:"devices"`
}

type Device struct {
	Name    string `json:"name"`
	Battery int    `json:"battery"`
}

/*
Run :- go run main.go < user.json
*/

func main() {
	var input []byte // reading file into bytes
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	if err := json.Unmarshal(input, &users); err != nil { // Converting bytes into structs using json.Unmarshal()
		fmt.Println(err) // Need to Pass address to User array
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}

		for _, device := range user.Devices {
			fmt.Printf("\n\t[ %-10s: %d%% ]", device.Name, device.Battery)
		}
		fmt.Println()
	}
}
