package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No character given")
		os.Exit(1)
	}

	starting_rune := rune(os.Args[1][0])
	if starting_rune < 'a' || starting_rune > 'z' {
		os.Exit(0)
	}

	var output strings.Builder
	for r := starting_rune; r <= 'z'; r++ {
		output.WriteRune(r)
	}
	fmt.Println(output.String())
}
