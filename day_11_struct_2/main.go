package main

import (
	"fmt"
	"time"
)

// Employee struct
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

// Point struct
type Point struct {
	X, Y int
}

// T struct
type T struct {
	x, y int
}

// Circle a circle
type Circle struct {
	Point
	Radius int
}

// Wheel a wheel
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	p := Point{1, 2}
	p1 := Point{X: 1}
	fmt.Println(p, p1)
	t := T{1, 2}
	fmt.Println(t)
	fmt.Println(Scale(p, 10))
	e := Employee{Salary: 20000000}
	fmt.Println(Bonus(&e, 10))
	AwardAnnualRaise(&e)
	fmt.Println(e.Salary)
	pp := &Point{1, 2}
	pp1 := new(Point)
	*pp1 = Point{1, 2}
	fmt.Println(pp, pp1)
	p2 := Point{1, 2}
	p3 := Point{1, 2}
	fmt.Println(p1 == p2, p2 == p3)
	var w Wheel
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Println(w)
}

// Scale scale a point with specific factor
func Scale(p Point, factor int) Point {
	pScale := Point{p.X * factor, p.Y * factor}
	return pScale
}

// Bonus caculcate bonus
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// AwardAnnualRaise func
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

// tree sort
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
