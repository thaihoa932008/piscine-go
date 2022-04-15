package main

import "fmt"

func Fibonacci2(index int) int {
	// the 1st 2 terms are 0 and 1, therefore when it's 0 and 1, it should return 0 and 1
	//all other terms are obtained by adding preceding two terms, therefore nth term is the sum of (n-1)th and (n-2)th terms
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci2(index-1) + Fibonacci2(index-2)
}

func main() {
	arg1 := 6
	fmt.Println(Fibonacci2(arg1))
}
