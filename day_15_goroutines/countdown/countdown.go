package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown ...")
	tick := time.Tick(1 * time.Second)
	for cd := 10; cd > 0; cd-- {
		fmt.Println(cd)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
