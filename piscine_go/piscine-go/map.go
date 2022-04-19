package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	// you create a var to store the slices of bool
	var result []bool
	for _, boolval := range a {
		result = append(result, f(boolval))
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime.a)
	fmt.Println(result)
}
