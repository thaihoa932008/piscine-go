package piscine

func IsNumeric(s string) bool {
	a := []rune(s)
	for _, num := range a {
		if !(num >= 48 && num <= 57) {
			return false
		}
	}
	return true
}
