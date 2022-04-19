package piscine

func Map(f func(int) bool, a []int) []bool {
	// you create a var to store the slices of bool
	var result []bool
	// you make a range loop to go through every value of the slice of a
	for _, boolval := range a {
		// the result go through the function of the loop and add the bool value to the result
		result = append(result, f(boolval))
	}
	return result
}
