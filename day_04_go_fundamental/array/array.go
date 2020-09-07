package array

import (
	"bytes"
	"fmt"
)

var buffer [256]byte
var x [2]string

// MyArray func
func MyArray() {
	for i, v := range buffer {
		fmt.Printf("byte %v: %v\n", i, v)
	}
	for i, v := range x {
		fmt.Printf("char %v: %v\n", i, v)
	}
}

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

// AddOneToEachElement func
func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

// SubtractOneFromLength func
func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

// MySlices func
func MySlices() {
	var slice1 []byte = buffer[100:150]
	var slice2 []byte
	slice2 = slice1[10:20]
	slice3 := buffer[10:20]
	sliceH1 := sliceHeader{
		Length:        50,
		ZerothElement: &buffer[100],
	}
	sliceH2 := sliceHeader{
		Length:        10,
		ZerothElement: &buffer[110],
	}
	slice1 = slice1[1 : len(slice1)-1]
	sliceH1Now := sliceHeader{
		Length:        48,
		ZerothElement: &buffer[101],
	}
	fmt.Printf("slice1 %v\n", slice1)
	fmt.Printf("slice2 %v\n", slice2)
	fmt.Printf("slice3 %v\n", slice3)
	fmt.Println(sliceH1)
	fmt.Println(sliceH1Now)
	fmt.Println(sliceH2)
}

// PassingSlicesToFunctions func
func PassingSlicesToFunctions() {
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println(buffer)
}

// PtrSubtractOneFromLength func
func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	fmt.Println(i)
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// PointerToSlice func
func PointerToSlice() {
	// slice := buffer[10:20]
	// fmt.Println("Before: len(slice) =", len(slice))
	// PtrSubtractOneFromLength(&slice)
	// fmt.Println("After:  len(slice) =", len(slice))
	// fmt.Println("After: slice =", slice)
	pathName := path("/usr/bin/tso")
	// fmt.Printf("%v\n", pathName)
	// pathName.TruncateAtFinalSlash()
	// fmt.Printf("%s\n", pathName)
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}
