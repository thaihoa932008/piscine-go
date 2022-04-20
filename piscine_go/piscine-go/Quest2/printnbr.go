// Write a function that prints an int passed in parameter. All possible values of type int have to go through. You cannot convert to int64

// 1st you declare a var is a slice of rune cause the result value is rune
// 2nd you declare the condition, if it's negative number, print the -, if it's = 0, print 0
// we then take the modulus of the integer with 10, and if the remain <0, then we turn it to regative nb by * with -1, and then add the remain back to the result
package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var result []rune
	if n < 0 {
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune(48)
	} else {
		for {
			var remain int = n % 10
			if remain < 0 {
				remain = remain * -1
			}
			result = append(result, rune(remain))
			n = n / 10
			if n == 0 {
				break
			}
		}
	}
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i] + 48)
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
