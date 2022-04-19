// Write a function CountIf that returns, the number of elements of a string slice, for which the f function returns true.

package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, result := range tab {
		if f(result) == true {
			count++
		}
	}
	return count
}
