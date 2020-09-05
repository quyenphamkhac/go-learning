package main

import (
	"fmt"
	"go-learning/day_02_variables_values_type/template"
)

func main() {
	var message string
	var a, b, c int
	a = 1
	var message1 = "Hello, world!"
	var d, e, f = 1, 2, 3
	message = "Hello, I'm Pham Khac Quyen"
	n, err := fmt.Println(message, a, b, c)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(n)
	}
	m, _ := fmt.Println(message1, d, e, f)
	fmt.Println(m)
	// Call modules
	template.Template()
	// Shorthand declaration
	message3 := "Shorthand declare"
	g, h, j := 4, 5, 6
	fmt.Println(message3, g, h, j)
}

func foo() {
	fmt.Println("I'm foo.")
}
