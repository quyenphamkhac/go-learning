package main

import (
	"flag"
	"fmt"
	"go-learning/day_18_interfaces/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
