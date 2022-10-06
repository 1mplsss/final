package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	a := args[0]
	b := args[1]

	if a == b {
		z01.PrintRune('0')
	} else if a < b {
		z01.PrintRune('-')
		z01.PrintRune('1')
	} else {
		z01.PrintRune('1')
	}
	z01.PrintRune('\n')
}
