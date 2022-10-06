package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := len(os.Args)
	if a > 2 || a == 0 {
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	if n == 0 {
		return
	}
	for {
		if n%2 == 0 {
			n = n / 2
			if n/2 == 1 {
				fmt.Println("true")
				break
			} else if n%2 == 1 {
				fmt.Println("false")
				break
			}
		}
	}
}
