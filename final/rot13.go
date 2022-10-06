package piscine

func Rot13(s string) string {
	res := []rune(s)
	for i := 0; i < len(s); i++ {
		if res[i] >= 'a' && res[i] <= 'm' || res[i] >= 'A' && res[i] <= 'M' {
			res[i] = res[i] + 13
		} else {
			res[i] = res[i] - 13
		}
	}
	// 	res := 0
	// for _, i := range s {
	// 	if i >= 'a' && i <= 'm' || i >= 'A' && i <= 'M' {
	// 		res[i] += 13		} else {
	// 		res = append(res, rune(i-13))
	// 	}
	// }
	return string(res)
}
