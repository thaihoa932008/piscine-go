package main

import (
	"fmt"
	"os"
)

func main() {
	str1 := "01"
	str2 := "galaxy"
	str3 := "galaxy 01"
	str := os.Args
	for _, i := range str {
		if i == str1 || i == str2 || i == str3 {
			fmt.Println("Alert!!!")
		}
	}
}
