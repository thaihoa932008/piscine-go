package main

import "fmt"

func ActiveBits(n int) int {
	var countbit int
	for n != 0 {
		countbit = countbit + n&1
	}
	return countbit
}

func main() {
	fmt.Println(ActiveBits(7))
}
