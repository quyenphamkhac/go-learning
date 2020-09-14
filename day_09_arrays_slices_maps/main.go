package main

import (
	"fmt"
	"sort"
)

func main() {
	// CompositeTypeArray()
	CompositeTypeMaps()
}

// CompositeTypeArray func
func CompositeTypeArray() {
	q := [...]int{1, 3, 4}
	fmt.Printf("%T  %v\n", q, q)
	q = [...]int{4, 5, 6}
	fmt.Printf("%T  %v\n", q, q)
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
}

// CompositeTypeMaps func
func CompositeTypeMaps() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	delete(ages, "alice")
	delete(ages1, "charlie")
	ages["sage"] = 53
	fmt.Println(ages, ages1)
	for k, v := range ages {
		fmt.Printf("%s \t%d\n", k, v)
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	var ages3 map[string]int
	ages3 = make(map[string]int)
	ages3["sage"] = 33
	age, ok := ages3["bob"]
	if !ok {
		fmt.Println("Not in map")
	} else {
		fmt.Println(age)
	}
	fmt.Println(ages3)
	// seen := make(map[string]bool)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	line := input.Text()
	// 	if !seen[line] {
	// 		seen[line] = true
	// 		fmt.Println(line)
	// 	}
	// }
}

// isMapEqual
func isMapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, vx := range x {
		if vy, ok := y[k]; !ok || vy != vx {
			return false
		}
	}
	return true
}
