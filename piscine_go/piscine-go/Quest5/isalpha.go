package piscine

func IsAlpha(s string) bool {
	if s == " " {
		return true
	}
	isA := []rune(s)
	for _, lens := range isA {
		if !(lens >= 65 && lens <= 90) && !(lens >= 97 && lens <= 122) && !(lens >= 48 && lens <= 57) {
			return false
		}
	}
	return true
}
