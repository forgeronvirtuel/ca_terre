package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Treating error
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Please provide (only) two numbers")
		os.Exit(1)
	}
	b, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while converting the base. Reason: %s", err)
		os.Exit(2)
	}
	e, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while converting the exponent. Reason: %s", err)
		os.Exit(3)
	}
	if e < 0 {
		fmt.Fprintln(os.Stderr, "Exponent cannot be negative")
		os.Exit(4)
	}

	var res int = 1
	for cptr := 1; cptr <= e; cptr++ {
		res *= b
	}

	fmt.Println(res)
}
