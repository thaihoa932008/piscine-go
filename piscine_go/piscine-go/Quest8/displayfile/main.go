package main

import (
	"fmt"
	"os"
)

func main() {
	// first you declare the arguments, start at the 1st position cause 0 is the name
	arg := os.Args[1:]
	// set the argument according to the requirements
	if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if len(arg) > 1 {
		fmt.Println("Too many arguments")
	} else {
		// follow the file manipulation in Go, GG more the file manipulation in Go
		file, err := os.Open("quest8.txt")
		arr := make([]byte, 14)
		file.Read(arr)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(arr))
		file.Close()
	}
}
