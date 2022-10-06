package piscine

func StrRev(s string) string {
	res := ""
	for _, i := range s {
		res = string(i) + res
	}
	return res
}
