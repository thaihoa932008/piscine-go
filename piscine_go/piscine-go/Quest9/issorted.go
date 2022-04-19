/* Write a function IsSorted that returns true, if the slice of int is sorted, otherwise returns false.

The function passed in the parameter returns a positive int if a (the first argument) is greater than to b (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.

To do your testing you have to write your own f function.
*/

package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + i; j < len(a); j++ {
			if AscendSorted(a) || DescendSorted(a) {
				return true
			}
		}
	}
	return false
}

func DescendSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				return false
			}
		}
	}
	return true
}

func AscendSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				return false
			}
		}
	}
	return true
}
