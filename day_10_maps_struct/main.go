package main

import (
	"fmt"
	"time"
)

func main() {
	// Map2()
	Struct()
}

// Employee struct
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

// Struct func
func Struct() {
	dilbert.Salary = 1000
	dilbert.Position = "Developer"
	pos := &dilbert.Position
	*pos = "Senior " + *pos
	fmt.Println(dilbert)
	var eotm *Employee = &dilbert
	eotm.Position += " (fuck yeah)"
	fmt.Println(dilbert, *eotm)
}

// Map2 func
func Map2() {
	addEdge("hello", "world")
	b := hasEdge("hello", "world")
	fmt.Println(b)
	var g = make(map[string]map[string]int)
	mm, ok := g["hello"]
	if !ok {
		mm = make(map[string]int)
		g["hello"] = mm
	}
	mm["world"]++
	fmt.Println(g)
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges, ok := graph[from]
	if !ok {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
