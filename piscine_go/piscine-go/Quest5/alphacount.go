package piscine

func AlphaCount(s string) int {
	a := []byte(s)
	count := 0
	for _, lens := range a {
		if lens >= 65 && lens <= 90 || lens >= 97 && lens <= 122 {
			count++
		}
	}
	return count
}
