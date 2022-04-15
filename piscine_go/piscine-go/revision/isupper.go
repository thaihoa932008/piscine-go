package main

import "fmt"

func IsUpper2(s string) bool {
	bytes := []byte(s)
	for _, result := range bytes {
		if !(result >= 65 && result <= 90) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper2("HELLO"))
	fmt.Println(IsUpper2("HELLO!"))

}
