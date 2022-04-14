package piscine

func ToUpper(s string) string {
	var newString string
	for _, str := range s {
		if str >= 'a' && str <= 'z' {
			newString += string(str - 32)
		} else {
			newString += string(str)
		}
	}
	return newString
}
