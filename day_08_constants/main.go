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

// MyFloat64 type
type MyFloat64 float64

// Zero constants
const Zero = 0.0

// TypedZero constants
const TypedZero float64 = 0.0

// Huge constants
const Huge = 1e1000

// MaxUint constants
const MaxUint = ^0

// Constants func
func Constants() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Printf("%T %[1]v\n", Sunday)
	var mf MyFloat64
	mf = 0.0
	mf = Zero
	fmt.Printf("%f\n", mf)
	fmt.Println(Huge / 1e999)
	pi := math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi, pi32)
	fmt.Println(MaxUint)
	var f float32 = 1
	var i int = 1.000
	var u uint32 = 1e3 - 99.0*10.0 - 9
	var c float64 = '\x01'
	var p uintptr = '\u0001'
	var r complex64 = 'b' - 'a'
	var b byte = 1.0 + 3i - 3.0i

	fmt.Println(f, i, u, c, p, r, b)

	var f1 = 'a' * 1.5
	fmt.Println(f1)
	// var i1 int = 2
	sqrt2 := math.Sqrt(2.0)
	fmt.Println(sqrt2)
}
