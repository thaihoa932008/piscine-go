package piscine

func AlphaCount(s string) int {
	//convert the string into byte
	a := []byte(s)
	count := 0
	//use the for range loop to go through the string of bytes
	for _, lens := range a {
		if lens >= 65 && lens <= 90 || lens >= 97 && lens <= 122 {
			// sort it out by the ascii value of the alphabet letter uppercase and lowercase
			count++
		}
	}
	return count
}
