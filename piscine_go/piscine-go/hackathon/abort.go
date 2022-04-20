package piscine

func Abort(a, b, c, d, e int) int {
	numslice := []int{a, b, c, d, e}

	for i := 0; i < len(numslice)-1; i++ {
		for j := i + 1; j < len(numslice); j++ {
			if numslice[i] > numslice[j] {
				numslice[i], numslice[j] = numslice[j], numslice[i]
			}
		}
	}
	return numslice[2]
}
