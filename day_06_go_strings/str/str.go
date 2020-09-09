package str

import (
	"fmt"
)

// Strings func
func Strings() {
	s := "hello, world"
	fmt.Println(&s)
	s = "fuck"
	var s1 string
	s1 = s
	fmt.Println(&s1)
	s = s + "1"
	fmt.Println(s1)
	const GoUsage = `Go is a tool for managing Go source code.
    Usage:
    	go command [arguments]
	...`
	fmt.Println(GoUsage)
}

// UnicodeShowCase func
func UnicodeShowCase() {

}
