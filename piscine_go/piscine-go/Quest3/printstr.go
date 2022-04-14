package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	slice := []rune(s)
	for _, a := range slice {
		z01.PrintRune(a)
	}
}
