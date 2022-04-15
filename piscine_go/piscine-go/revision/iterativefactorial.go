package main

import "fmt"

func IterativeFactorial2(nb int) int {
	result := 1
	if nb < 0 || nb > 63 {
		return 0
	}
	if nb > 0 && nb < 63 {
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial2(arg))
}
