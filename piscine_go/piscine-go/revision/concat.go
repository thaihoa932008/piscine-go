package main

import "fmt"

func Concat(str1 string, str2 string) string {
	concats := str1 + str2
	return concats
}

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}
