package main

import "fmt"

func RecursiveFactorial2(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		return 0
	}
	if nb > 0 && nb < 63 {
		result = nb * RecursiveFactorial2(nb-1)
	}
	return result
}
func main() {
	arg := 4
	fmt.Println(RecursiveFactorial2(arg))
}
