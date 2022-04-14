package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for i := 2; i < nb/2+1; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
