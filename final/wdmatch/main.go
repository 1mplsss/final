package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	j := 0
	count := 0
	for i := range s2 {
		if j >= len(s1) {
			break
		}
		if s2[i] == s1[j] {
			j++
			count++
		}
	}
	if count == len(s1) {
		for _, ch := range s1 {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')

	} else {
		return
	}
}
