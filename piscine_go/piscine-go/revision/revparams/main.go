package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	slice := os.Args
	for i := len(slice) - 1; i > 0; i-- {
		for _, rev := range slice[i] {
			z01.PrintRune(rev)
		}
		z01.PrintRune('\n')
	}
}
