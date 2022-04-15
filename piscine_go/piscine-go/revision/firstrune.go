package main

import "github.com/01-edu/z01"

func FirstRune2(s string) rune {
	runes := []rune(s)
	return runes[0]
}

func main() {
	z01.PrintRune(FirstRune2("Hello!"))
	z01.PrintRune(FirstRune2("Salut!"))
	z01.PrintRune(FirstRune2("Ola!"))
	z01.PrintRune('\n')
}
