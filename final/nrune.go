package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	if len(s)-1 >= n && n > 0 {
		return a[n-1]
	}
	return 0
}
