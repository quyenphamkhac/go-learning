package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	names := []string{"Pham", "Khac", "Quyen"}
	var wg sync.WaitGroup
	ch := make(chan string)
	for _, n := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			res, err := printName(name)
			if err != nil {
				log.Println(err)
				return
			}
			ch <- res
		}(n)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var fullName string
	for s := range ch { //blocking
		fullName += s
	}
	fmt.Println("Full Name: ", fullName)
}

func printName(name string) (string, error) {
	time.Sleep(1 * time.Second)
	fmt.Println(name)
	return name, nil
}
