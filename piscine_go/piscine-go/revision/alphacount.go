package main

import "fmt"

func AlphaCount2(s string) int {
	count := 0
	runes := []rune(s)
	for _, rangerune := range runes {
		if (rangerune >= 65 && rangerune <= 90) || (rangerune >= 97 && rangerune <= 122) {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount2(s)
	fmt.Println(nb)
}
