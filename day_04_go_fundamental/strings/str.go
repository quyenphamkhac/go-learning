package strings

import "fmt"

// GoString func
func GoString() {
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
