package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	ex_prev := 0
	prev := 0
	count := 0
	return func() int {
		current := 0

		if count < 2 {
			current = count
		} else {
			current = ex_prev + prev
		}
		count += 1
		ex_prev = prev
		prev = current
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
