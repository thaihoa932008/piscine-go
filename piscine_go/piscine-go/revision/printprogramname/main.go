package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// create a slice that contain the argument, to print the name the argument is 0
	slices := os.Args[0]
	// convert it into runes
	runeS := []rune(slices)
	// run a loop to go through the rune, start from the 2nd position
	for _, name := range runeS[2:] {
		z01.PrintRune(name)
	}
	z01.PrintRune('\n')
}
