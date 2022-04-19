package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println("Program path:", os.Args[0])
	programname := path.Base(os.Args[0])
	fmt.Println("Program name:", programname)
}
