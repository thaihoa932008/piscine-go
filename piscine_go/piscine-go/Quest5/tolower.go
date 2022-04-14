package piscine

func ToLower(s string) string {
	var newString string
	for _, str := range s {
		if str >= 'A' && str <= 'Z' {
			// you add 32 to the value because the value of uppercase and lowercase character in the ASCII is parallel 32 different
			newString += string(str + 32)
		} else {
			newString += string(str)
		}
	}
	return newString
}
