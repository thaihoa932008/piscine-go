package piscine

func IsLower(s string) bool {
	sRune := []byte(s)
	for _, lens := range sRune {
		if !(lens >= 97 && lens <= 122) {
			/** with (len >=97 && lens <=122) the loop only check for the 1st letter
			if it's uppercase, it will return TRUE immediately **/
			/** that's why you want to use the !-reverse for the loop to run through
			the whole loop first, till the end to find any value that's not uppercase,
			and then return FALSE, with that you can check the whole string **/
			return false
		}
	}
	return true
}
