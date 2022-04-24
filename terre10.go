package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Checking arg format
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide an argument to the command line")
		os.Exit(1)
	}

	// Create the regexp and test if the format is correct
	validTime := regexp.MustCompile(`^(?P<hour>\d{2}):(?P<minute>\d{2})$`)
	if !validTime.MatchString(os.Args[1]) {
		fmt.Fprintf(os.Stderr, "Time %s does not match format: hh:mm\n", os.Args[1])
		os.Exit(2)
	}

	// Retrieve groups from the regexp matching
	match := validTime.FindStringSubmatch(os.Args[1])
	params := make(map[string]string)
	for i, name := range validTime.SubexpNames() {
		params[name] = match[i]
	}

	// Checking that there is no weird time
	h, _ := strconv.Atoi(params["hour"])
	m, _ := strconv.Atoi(params["minute"])
	if h < 0 || 23 < h {
		fmt.Fprintf(os.Stderr, "Error, hours should be 00 <= hours <= 23, got %d\n", h)
		os.Exit(3)
	}
	if m < 0 || 59 < m {
		fmt.Fprintf(os.Stderr, "Error, minutes should be 00 <= minutes <= 59 , got %d\n", m)
		os.Exit(4)
	}

	// Converting date formats
	if h == 0 {
		fmt.Printf("12:%sPM\n", params["minute"])
		return
	}
	if h <= 12 {
		fmt.Printf("%sAM\n", os.Args[1])
		return
	}
	fmt.Printf("%02d:%sPM\n", h-12, params["minute"])
}
