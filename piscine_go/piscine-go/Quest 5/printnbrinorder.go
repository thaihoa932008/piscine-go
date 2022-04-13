package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var arrayInt []int
	for n > 0 {
		arrayInt = append(arrayInt, n%10)
		n = n / 10
	}
	for i := 0; i < len(arrayInt)-1; i++ {
		for j := 0; j < len(arrayInt)-i-1; j++ {
			if arrayInt[j] > arrayInt[j+1] {
				arrayInt[j], arrayInt[j+1] = arrayInt[j+1], arrayInt[j]
			}
		}
	}
	for _, results := range arrayInt {
		z01.PrintRune(rune(results + '0'))
	}
}
