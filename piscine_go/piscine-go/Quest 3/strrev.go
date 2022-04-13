package piscine

func StrRev(s string) string {
	a := []rune(s)
	var b []rune
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return string(b)
}
