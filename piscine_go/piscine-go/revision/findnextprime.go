package main

import "fmt"

func FindNextPrime2(nb int) int {
	// here because the 1st prime nb is 2, so any nb below 2 should return 2
	if nb < 2 {
		return 2
	}
	//ask to explain the condition of i in the loop
	for i := 2; i < nb/2+1; i++ {
		//if the nb is not the prime, return the next nb, which position +1 in the loop
		if nb%i == 0 {
			return FindNextPrime2(nb + 1)
		}
	}
	return nb
}

func main() {
	fmt.Println(FindNextPrime2(5))
	fmt.Println(FindNextPrime2(4))
}
