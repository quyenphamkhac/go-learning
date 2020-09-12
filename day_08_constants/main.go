package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	Constants()
}

const (
	a = 1
	b
	c = 2
	d
)

// Weekday custom type
type Weekday int

// Weekday const define
const (
	Sunday Weekday = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Constants func
func Constants() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Printf("%T %[1]v\n", Sunday)
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}
