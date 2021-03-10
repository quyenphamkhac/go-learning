package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Strings func
func Strings() {
	s := "Hello, カタカナaaa"
	for _, r := range s {
		fmt.Printf("%#U\n", r)
	}
	s1 := "カタカナ"
	fmt.Printf("% x\n", s1)
	r := []rune(s1)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
}

// StringsByteSlices func
func StringsByteSlices() {
	s := "hello, ờ"
	rs := []rune(s)
	for _, r := range rs {
		fmt.Printf("Upper: %q %q\n", r, unicode.ToUpper(r))
	}
	s1 := "/a/b/c.go"
	fmt.Println(basename(s1))
	fmt.Println(basenameV2(s1))
	s2 := "123456789"
	fmt.Println(comma(s2))
	fmt.Println(commaV2(s2))
	s3 := "hello, world"
	b := []byte(s3)
	s4 := string(b)
	fmt.Println(s4)
	ints := []int{1, 2, 3, 4}
	s5 := intsToString(ints)
	fmt.Println(s5)
}

// commaV2 func
func commaV2(s string) string {
	if len(s) < 3 {
		return s
	}
	var buf bytes.Buffer
	d, t := len(s)%3, len(s)/3
	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}
	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d : i*3+d+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

// intsToString func
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// basename func
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[0:i]
			break
		}
	}
	return s
}

// basenameV2 func
func basenameV2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// strings = slice of bytes but immutable

// StringToNumber func
func StringToNumber() {
	// number to string
	x := 1234
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	x, err := strconv.Atoi("124a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}

func main() {
	// Strings()
	// StringsByteSlices()
	StringToNumber()
}
