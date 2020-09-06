package main

import (
	"fmt"
)

func main() {
	var x int8
	var y byte
	y = 199
	x = 126
	fmt.Println(x, y)
	s := "𩸽"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println()
	for _, v := range s {
		fmt.Println(v)
	}
}
