package main

import (
	"fmt"
)

func main() {
	fib := fibonacci()
	var sum int
	for f := 0; f < 4000000; f = fib() {
		if f%2 == 0 {
			sum += f
		}
	}
	fmt.Println(sum)
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
