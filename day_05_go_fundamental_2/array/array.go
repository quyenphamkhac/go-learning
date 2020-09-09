package array

import (
	"fmt"
)

var buffer [256]byte

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
		slice2 = Extend(slice2, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
		fmt.Println("address of 0th element:", &slice2[0])
	}
	var arr = [3]int{1, 3, 4}
	slice3 := arr[0:1]
	fmt.Println(arr, slice3)
	for i := 0; i < 10; i++ {
		slice3 = Extend(slice3, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
		fmt.Println("address of 0th element:", &slice3[0])
	}
	fmt.Println(arr)
}

// Insert func
func Insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

// ExtendAppendShowCase func
func ExtendAppendShowCase() {
	a := []int{1, 3, 4, 5}
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
}

// Extend func
func Extend(slice []int, element int) []int {
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

// Append func
func Append(slice []int, elements ...int) []int {
	// for _, item := range items {
	// 	slice = Extend(slice, item)
	// }
	// return slice
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

// NilSliceShowCase func
func NilSliceShowCase() {
	nilSlice := sliceHeader{
		Length:        0,
		Capacity:      0,
		ZerothElement: nil,
	}
	fmt.Println(nilSlice)
}
