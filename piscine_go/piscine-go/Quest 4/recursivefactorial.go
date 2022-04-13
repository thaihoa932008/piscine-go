package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		return 0
	}
	if nb > 0 && nb < 63 {
		return nb * RecursiveFactorial(nb-1)
	}
	return result
}
