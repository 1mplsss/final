package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	b := os.Args[1]
	runes := []rune(b)

	if len(a) == 1 {
		for _, i := range runes {
			if i >= 'Z' && i <= 'A' {
				z01.PrintRune('Z' - (i - 'A'))
			} else if i >= 'z' && i <= 'a' {
				z01.PrintRune('z' - (i - 'a'))
			}
		}
		z01.PrintRune('\n')
	}
}
