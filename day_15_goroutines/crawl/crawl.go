package main

import (
	"fmt"
	"go-learning/day_15_goroutines/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	<-tokens
	return list
}
