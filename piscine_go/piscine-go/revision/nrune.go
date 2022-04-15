package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	runes := []rune(s)
	// n is the position extract, when n <0, n is not in the string, so return 0
	if n < 0 {
		return 0
	}
	// when n > length of the string of rune, n is not in the string, so return 0
	if n > len(runes) {
		return 0
	}
	// else return nth-1 like in fibonacci
	return runes[n-1]
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
