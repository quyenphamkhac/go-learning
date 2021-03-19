package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	x := []string{"Pham", "Khac", "Quyen"}
	words := make(chan string)
	done := make(chan []string)
	go func() {
		for _, w := range x {
			words <- w
		}
		close(words)
	}()
	go func() {
		r := upper(words)
		done <- r
	}()
	fmt.Println(<-done)
}

func upper(words <-chan string) []string {
	var wg sync.WaitGroup
	uppers := make(chan string)
	for w := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			uppers <- strings.ToUpper(w)
		}(w)
	}
	go func() {
		wg.Wait()
		close(uppers)
	}()
	var result []string
	for w := range uppers {
		result = append(result, w)
	}
	return result
}
