package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := len(os.Args)
	z01.PrintRune((rune(args-1) + 48))
	z01.PrintRune('\n')
}
