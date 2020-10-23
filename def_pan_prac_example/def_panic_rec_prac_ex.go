package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Uing Defer keyword for closing resource ,
	// Defer keyword makes sure it is executed in last
	// Without Defer it would compulsory to close it in last and programmers may miss it
	roboots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", roboots)
}
