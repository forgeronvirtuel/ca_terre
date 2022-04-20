package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Need one string as arguments")
		os.Exit(1)
	}

	toInvert := os.Args[1]
	var builder strings.Builder
	for idx := len(toInvert) - 1; idx >= 0; idx-- {
		builder.WriteByte(toInvert[idx])
	}

	fmt.Println(builder.String())
}
