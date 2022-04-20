// 1st you create a new var will hold the value of new string
// 2nd run a range loop through the string and set condition
// condition
// z becomes n so from n to z, to change we -12
// later count from m to a, a change to o, so 14 value if diff, so to change it, we add 14
// last step, any other cases, don't change anything

package piscine

func Rot14(s string) string {
	var newStr string
	for _, str := range s {
		if str >= 'A' && str <= 'M' || str >= 'a' && str <= 'm' {
			newStr += string(str + 14)
		} else if str >= 'N' && str <= 'Z' || str >= 'n' && str <= 'z' {
			newStr += string(str - 12)
		} else {
			newStr += string(str)
		}
	}
	return newStr
}
