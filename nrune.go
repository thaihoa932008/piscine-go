package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 {
		return 0
	}
	if n > len(s) {
		return 0
	}
	return runes[n-1]
}
