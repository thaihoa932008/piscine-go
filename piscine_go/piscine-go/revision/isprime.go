package main

import "fmt"

func IsPrime2(nb int) bool {
	//prime numbers are numbers which are completely divisible by only themselves.
	//Completely divisible means if the number is divided by any of the number other than itself, the remainder of the division should be 0
	if nb == 1 {
		return false
	}
	if nb < 0 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime2(5))
	fmt.Println(IsPrime2(4))
}
