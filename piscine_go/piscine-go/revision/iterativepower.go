package main

import "fmt"

func IterativePower2(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		result = result * nb
	}
	return result
}
func main() {
	fmt.Println(IterativePower2(4, 3))
}
