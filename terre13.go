package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// last contain the value previously seen
	var last int
	// lastIsSet change of value when last is set for the first time
	var lastIsSet bool

	// Check if the list of number given in args is sorted
	for idx, arg := range os.Args[1:] {

		// Convert and check if argument is a number
		v, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot convert argument %s at position %d. Reason : %s\n", arg, idx, err)
			os.Exit(1)
		}

		// If last is not set, take the first arg
		if !lastIsSet {
			lastIsSet = true
			last = v
			continue
		}

		// If the current value is not superior to the last one, then it is not sorted
		if v < last {
			fmt.Fprintf(os.Stderr, "Value no %d is not sorted, got : %d, last was %d\n", idx+1, v, last)
			return
		}
		last = v
	}
	fmt.Println("Sorted.")
}
