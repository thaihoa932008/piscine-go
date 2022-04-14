package piscine

func Capitalize(s string) string {
	result := ""
	caps := false
	for i, letter := range s {
		a := s[i]
		if IsAlpha(string(a)) {
			if caps {
				if IsLower(string(a)) {
					result += string(a - 32)
					caps = false
				} else {
					result += string(letter)
					caps = false
				}
			} else if i == 0 {
				if IsLower(string(a)) {
					result += string(a - 32)
				} else {
					result += string(letter)
				}
			} else if IsUpper(string(a)) {
				result += string(a + 32)
			} else {
				result += string(letter)
			}
		} else {
			result += string(letter)
			caps = true
		}
	}
	return result
}
