package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*

go build -> gives executable named as folder name websites
./webstite < log.txt

*/

func main() {
	m := newmetrics()                // Get a  pointer to Metrics
	in := bufio.NewScanner(os.Stdin) // Getting input from Standard input

	for in.Scan() { // it scans line by line
		// Get Data from Line
		data := parse(m, in.Text()) // in.Text() lets you get the last line read
		update(m, data)             // update metrics based on data for a domain
	}

	summarise(m) // Print the Summary
	// in.Err() -> gives the Error in processing Standard input
	dumpErrs([]error{in.Err(), err(m)})
}

// Printing  Summary
func summarise(m *metrics) {
	sort.Strings(m.domains)

	fmt.Printf("%-30s  %10s\n", "DOMAIN", "VIEWS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range m.domains {
		fmt.Printf("%-30s  %10d\n", domain, m.stats[domain].visits)
	}
	fmt.Printf("\n%-30s  %10d\n", "TOTAL", m.lines)
}

func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println(" Err:", err)
		}
	}
}
