package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Need only one argument")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert argument to int. Reason: %s\n", err)
		os.Exit(2)
	}

	if n < 0 {
		fmt.Printf("Cannot test negative number, got %d\n", n)
		return
	}

	if n == 1 || n == 2 {
		fmt.Printf("%d is not a prime number\n", n)
		return
	}

	for d := 2; d < n; d++ {
		if n%d == 0 {
			fmt.Printf("%d is not a prime number\n", n)
			return
		}
	}
	fmt.Printf("%d is a prime number\n", n)
}
