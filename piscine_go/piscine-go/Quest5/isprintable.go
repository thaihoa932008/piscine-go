package piscine

func IsPrintable(s string) bool {
	a := []rune(s)
	for _, printable := range a {
		if !(printable >= 32 && printable <= 127) {
			return false
		}
	}
	return true
}
