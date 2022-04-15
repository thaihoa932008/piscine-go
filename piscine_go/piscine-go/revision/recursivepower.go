package main

import "fmt"

func RecursivePower2(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power > 0 {
		result = nb * RecursivePower2(nb, power-1)
	}
	return result
}
func main() {
	fmt.Println(RecursivePower2(4, 3))
}
