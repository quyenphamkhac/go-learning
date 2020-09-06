package main

import (
	"fmt"
)

type horobi string

func main() {
	// ex1
	x, y, z := 42, "James Bond", true
	fmt.Println(x, y, z)
	var x1 int
	var y1 string
	var z1 bool
	// zero value
	fmt.Println(x1, y1, z1)
	x1 = 43
	y1 = "Bond James"
	z1 = false
	s := fmt.Sprintf("%v%v%v", x1, y1, z1)
	fmt.Println(s)
	var jin horobi
	jin = "Son"
	fmt.Printf("%v - %T", jin, jin)
}
