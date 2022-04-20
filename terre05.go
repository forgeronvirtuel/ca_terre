package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Need two numbers to compute the division")
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing the dividend. Reason: %s", err)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing the divider. Reason: %s", err)
	}

	if b == 0 {
		fmt.Fprintln(os.Stderr, "Cannot divide by 0")
		os.Exit(1)
	}

	fmt.Printf("Quotien: %d, Remainder: %d\n", a/b, a%b)
}
