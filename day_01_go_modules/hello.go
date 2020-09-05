package hello

import "rsc.io/quote/v3"

// Hello func
func Hello() string {
	return quote.HelloV3()
}

// Proverb func
func Proverb() string {
	return quote.Concurrency()
}
