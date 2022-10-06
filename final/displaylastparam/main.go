package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, i := range args[len(args)-1] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
