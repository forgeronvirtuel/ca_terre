package main

import (
	"fmt"
	"math"
	"os"
)

func isNotDigit(r byte) bool {
	return r < '0' || '9' < r
}

func isDigit(r byte) bool {
	return '0' <= r && r <= '9'
}

func atoi(s string) (int, error) {
	var v, ctr int

	// Loop over each rune to check if it is a number and to add it
	// to the whole converted number
	for idx := len(s) - 1; idx > 0; idx, ctr = idx-1, ctr+1 {
		r := s[idx]
		if isNotDigit(r) {
			return 0, fmt.Errorf("%s (char %d) is not convertible to int", s, idx)
		}
		v += int(r-'0') * int(math.Pow10(ctr))
	}

	frstRune := s[0]
	if frstRune == '-' {
		v = -v
	} else if isDigit(frstRune) {
		v += int(frstRune) * int(math.Pow10(ctr+1))
	} else if frstRune != '+' {
		return 0, fmt.Errorf("%s (char 0: %c) is not convertible to int", s, frstRune)
	}

	return v, nil
}

func main() {

	// Test if there is one and only one argument from the command line
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Give only one number to test on the command line")
		os.Exit(1)
	}

	// Try to convert the given argument to an integer
	arg := os.Args[1]
	// v, err := strconv.Atoi(arg)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Cannot convert %s to int. Reason: %s\n", arg, err)
	// 	os.Exit(2)
	// }
	v, err := atoi(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot convert %s to int. Reason: %s\n", arg, err)
		os.Exit(2)
	}

	// Test if the argument is even or odd
	if v%2 == 0 {
		fmt.Println("Pair")
	} else {
		fmt.Println("Impair")
	}
}
