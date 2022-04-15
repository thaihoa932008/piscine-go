package main

import "fmt"

func IsPrintable(s string) bool {
	bytes := []rune(s)
	for _, result := range bytes {
		if !(result >= 32 && result <= 127) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
