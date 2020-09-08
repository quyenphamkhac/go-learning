package array

import (
	"fmt"
)

var buffer [256]byte

// Extend func
func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

type sliceHeader struct {
	Length        int
	Capacity      int
	ZerothElement *int
}

// CapacityShowCase func
func CapacityShowCase() {
	var iBuffer [10]int
	slice := iBuffer[1:3]
	sliceH1 := sliceHeader{
		Length:        1,
		Capacity:      7,
		ZerothElement: &iBuffer[0],
	}
	for i := 0; i < 7; i++ {
		slice = Extend(slice, i)
		fmt.Println(slice)
		if len(slice) == cap(slice) {
			fmt.Println("slice is full !!")
		}
	}
	fmt.Println(sliceH1)
}

// Gopher type
type Gopher int

// MakeArrayShowCase func
func MakeArrayShowCase() {
	slice := make([]int, 10, 15)
	for i := range slice {
		slice[i] = i
	}
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	// manual copy
	for i := range slice {
		newSlice[i] = slice[i]
	}
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	fmt.Println(slice)
}

// CopySliceShowCase func
func CopySliceShowCase() {
	slice := make([]int, 10, 15)
	for i := range slice {
		slice[i] = i
	}
	newSlice := make([]int, len(slice), 2*cap(slice))
	copy(newSlice, slice)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	insertedSlice := Insert(newSlice, 10, 100)
	fmt.Println(insertedSlice, len(insertedSlice), cap(insertedSlice))
	slice1 := make([]int, 10, 20)
	for i := range slice1 {
		slice1[i] = i
	}
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = Insert(slice1, 5, 99)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice2 := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice2 = ExtendV2(slice2, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
		fmt.Println("address of 0th element:", &slice2[0])
	}
}

// Insert func
func Insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

// ExtendV2 func
func ExtendV2(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
