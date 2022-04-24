package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Checking args numbers
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide an argument to the command line")
		os.Exit(1)
	}

	// Create the regexp and test if the format is correct
	validTime := regexp.MustCompile(`^(?P<hour>\d{2}):(?P<minute>\d{2})(?P<indice>AM|PM)$`)
	if !validTime.MatchString(os.Args[1]) {
		fmt.Fprintf(os.Stderr, "Time %s does not match format: hh:mm<AM|PM>\n", os.Args[1])
		os.Exit(2)
	}

	// Retrieving groups
	match := validTime.FindStringSubmatch(os.Args[1])
	params := make(map[string]string)
	for i, name := range validTime.SubexpNames() {
		params[name] = match[i]
	}

	// Checking that there is no weird time
	h, _ := strconv.Atoi(params["hour"])
	m, _ := strconv.Atoi(params["minute"])
	if h < 1 || 12 < h {
		fmt.Fprintf(os.Stderr, "Error, hours should be 01 <= hours <= 12, got %d\n", h)
		os.Exit(3)
	}
	if m < 0 || 59 < m {
		fmt.Fprintf(os.Stderr, "Error, minutes should be 00 <= minutes <= 59 , got %d\n", m)
		os.Exit(4)
	}

	// if AM, print directly
	ind := params["indice"]
	if ind == "AM" {
		fmt.Printf("%02d:%02d\n", h, m)
	}
	// if PM and h=12 => h=00
	if ind == "PM" && h == 12 {
		fmt.Printf("00:%02d\n", m)
		return
	}
	// if PM, h += 12
	if ind == "PM" {
		fmt.Printf("%02d:%02d\n", h+12, m)
	}
}
