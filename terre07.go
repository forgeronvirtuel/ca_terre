package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Treating error case
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error, need only one string as argument")
	}
	if _, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Fprintln(os.Stderr, "Error, please provide a string not a number")
	}

	fmt.Println(len(os.Args[1]))
}
