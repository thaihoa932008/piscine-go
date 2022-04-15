package main

import "fmt"

func IsAlpha2(s string) bool {
	bytes := []rune(s)
	for _, results := range bytes {
		if !(results >= 65 && results <= 90) && !(results >= 97 && results <= 122) && !(results >= 48 && results <= 57) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha2("Hello! How are you?"))
	fmt.Println(IsAlpha2("HelloHowareyou"))
	fmt.Println(IsAlpha2("What's this 4?"))
	fmt.Println(IsAlpha2("Whatsthis4"))

}
