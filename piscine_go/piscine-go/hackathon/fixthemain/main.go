package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

type Door struct {
	state bool
}

var OPEN = false
var CLOSE = true

func OpenDoor(ptrDoor *Door) {
	fmt.Println("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	fmt.Println("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor *Door) bool {
	fmt.Println("Is the Door Opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	fmt.Println("Is the Door Closed ?")
	return ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)

	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
