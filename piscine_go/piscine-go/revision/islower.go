package main

import "fmt"

func IsLower(s string) bool {
	bytes := []byte(s)
	for _, result := range bytes {
		if !(result >= 97 && result <= 122) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
