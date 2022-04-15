package main

import "fmt"

func ToUpper2(s string) string {
	var newStr string
	bytes := []byte(s)
	for _, str := range bytes {
		if str >= 97 && str <= 122 {
			newStr = newStr + string(str-32)
		} else {
			newStr = newStr + string(str)
		}
	}
	return newStr
}

func main() {
	fmt.Println(ToUpper2("Hello! How are you?"))
}
