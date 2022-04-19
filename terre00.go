package main

import (
	"fmt"
	"strings"
)

func main() {
	var output strings.Builder
	for c := 'a'; c <= 'z'; c++ {
		output.WriteRune(c)
	}
	fmt.Println(output.String())
}
