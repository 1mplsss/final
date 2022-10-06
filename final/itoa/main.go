package piscine

func Itoa(n int) string {
	var arrayofRunes []rune
	s := ""
	if n < 0 {
		for n != 0 {
			arrayofRunes = append(arrayofRunes, rune(-(n%10)+48))
			n /= 10
		}
	} else if n > 0 {
		for n != 0 {
			arrayofRunes = append(arrayofRunes, rune(n%10+48))
			n /= 10
		}
	} else {
		arrayofRunes = append(arrayofRunes, '0')
	}

	for i := len(arrayofRunes) - 1; i >= 0; i-- {
		s += string(arrayofRunes[i])
	}
	return s
}
