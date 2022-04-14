package piscine

func TrimAtoi(s string) int {
	num := 0
	negative := false
	// create a range loop that go through the string
	for _, tempNum := range s {
		// the tempNum looks for ascii value inside this range of all the numeric
		if tempNum >= 48 && tempNum <= 57 {
			// when the number is found, like 0 for exp, 0*10 + 1 = 1; 1*10 + 2 = 12
			num = (num * 10) + (int(tempNum) - 48)
		}
		if num <= 0 {
			if tempNum == 45 {
				//if the value is -
				negative = true
			}
		}
	}
	if negative {
		//and when negative is true, return the number * -1 which turn it into the negative number
		// otherwise, return the number
		return (num * -1)
	}
	return num
}
