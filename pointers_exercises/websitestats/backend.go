package main

import (
	"fmt"
	"strconv"
	"strings"
)

type data struct {
	domain string
	visits int
}

type metrics struct {
	stats   map[string]data // stats per domian
	domains []string
	total   int   // Total visits
	lines   int   // lines processed
	lerr    error // Any error intercepted so far
}

func newmetrics() *metrics {
	return &metrics{stats: make(map[string]data)}
}

/*
Notice the Return type's name and how its used in the snippet below
Passing metrics as pointer since it has to edited
*/
func parse(m *metrics, line string) (d data) { // Return variable d is already declared like this
	if m.lerr != nil { // Error encountered no need to process further
		return
	}
	fields := strings.Fields(line) // Splits Line based on spaces

	m.lines++
	if len(fields) != 2 { // each line must contains domain and views it received
		m.lerr = fmt.Errorf("incorrect format %v  at line %d ", fields, m.lines)
		return
	}

	var err error
	d.domain = fields[0]
	d.visits, err = strconv.Atoi(fields[1])

	if d.visits < 0 || err != nil {
		m.lerr = fmt.Errorf("wrong input %q at line %d ", d.visits, err)
	}
	return // no need to type d as well since it has been informed in Return signature of Function
}

func update(m *metrics, d data) {
	if m.lerr != nil {
		return
	}

	// Check for unique domains
	// Map returns false in case the key is missing
	if _, ok := m.stats[d.domain]; !ok {
		m.domains = append(m.domains, d.domain) // Add the new domain
	}
	m.total += d.visits

	// creating / updating stats for domain
	m.stats[d.domain] = data{
		domain: d.domain,
		visits: d.visits + m.stats[d.domain].visits,
	}
}

// function to return Error encountered
func err(m *metrics) error {
	return m.lerr
}
