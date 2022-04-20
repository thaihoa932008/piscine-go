package piscine

func Unmatch(a []int) int {
	for _, i := range a {
		var res int
		for _, j := range a {
			if i == j {
				res++
			}
		}
		if res%2 == 1 || res == 1 {
			return i
		}
	}
	return -1
}
