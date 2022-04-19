// Write a function Any that returns true, for a string slice :
// -if, when that string slice is passed through an f function, at least one element returns true.

package piscine

func Any(f func(string) bool, a []string) bool {
	for _, result := range a {
		if f(result) {
			return true
		}
	}
	return false
}
