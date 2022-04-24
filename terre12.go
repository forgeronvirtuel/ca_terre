package main

import (
	"fmt"
	"os"
	"strconv"
)

func getMiddleOfThree(a, b, c int) int {
	if (b <= a && a <= c) || (c <= a && a <= b) {
		return a
	}
	if (a <= b && b <= c) || (c <= b && b <= a) {
		return b
	}
	return c
}

func main() {
	// Checking that there is 3 args
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Provide three numbers as arguments")
		os.Exit(1)
	}

	// Converting
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert first argument to int. Reason : %s\n", err)
		os.Exit(2)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert second argument to int. Reason : %s\n", err)
		os.Exit(3)
	}
	c, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert third argument to int. Reason : %s\n", err)
		os.Exit(4)
	}

	// if 3 args are identical => error
	if a == b && b == c {
		fmt.Fprintln(os.Stderr, "All arguments are equals")
		os.Exit(5)
	}

	// Get the middle
	fmt.Println(getMiddleOfThree(a, b, c))
}
