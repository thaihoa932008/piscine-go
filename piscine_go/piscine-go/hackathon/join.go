package piscine

func Join(strs []string, sep string) string {
	var newStr string
	for i, str := range strs {
		newStr += str
		if i != len(strs)-1 {
			newStr += sep
		}
	}
	return newStr
}
