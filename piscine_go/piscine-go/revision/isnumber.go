package main

import "fmt"

func IsNumeric(s string) bool {
	bytes := []rune(s)
	for _, result := range bytes {
		if !(result >= 48 && result <= 57) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
