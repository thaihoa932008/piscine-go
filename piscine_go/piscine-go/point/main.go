package main

/* with x and y as rune
type point struct {
	x []rune
	y []rune
}

func setPoint(ptr *point) {
	ptr.x = []rune("42")
	ptr.y = []rune("21")
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintStr(string(points.x))
	PrintStr(",")
	PrintStr(" ")
	PrintStr("y = ")
	PrintStr(string(points.y))
	z01.PrintRune('\n')
}  */

/* with x and y as int
type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(s string) {
	for _, item := range s {
		z01.PrintRune(item)
	}
}

func PrintNbrInOrder(n int) {
	nb := 0
	var array []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		nb = n % 10
		n = n / 10
		array = append(array, nb)
	}
	for i := len(array) - 1; i >= 0; i-- {
		z01.PrintRune(rune(array[i] + 48))
	}
}

func main() {
	points := &point{}
	str1 := "x = "
	str2 := ", y = "
	setPoint(points)
	PrintStr(str1)
	PrintNbrInOrder(points.x)
	PrintStr(str2)
	PrintNbrInOrder(points.y)
	z01.PrintRune('\n')
} */

/* with x and y as string
type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}
	setPoint(points)
	str := "x = " + points.x + ", y = " + points.y
	PrintStr(str)
}


*/
