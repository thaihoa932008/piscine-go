package main

import "fmt"

func ToLower2(s string) string {
	var newStr string
	bytes := []byte(s)
	for _, str := range bytes {
		if str >= 65 && str <= 90 {
			newStr = newStr + string(str+32)
		} else {
			newStr = newStr + string(str)
		}
	}
	return newStr
}

func main() {
	fmt.Println(ToLower2("Hello! How are you?"))
}
