package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	slices := os.Args
	for i := len(slices) - 1; i > 0; i-- {
		// the first loop is used to reverse the order, count from the last
		// the 2nd loop is used to print the slices in runes
		for _, space := range slices[i] {
			z01.PrintRune(space)
		}
		z01.PrintRune('\n')
	}
}
