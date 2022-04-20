package piscine

func ActiveBits(n int) int {
	var countbit1 int
	for n > 0 {
		if n == 0 {
			return 0
		}
		n &= (n - 1)
		countbit1++
	}
	return countbit1
}
