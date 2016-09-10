package main

import "fmt"

func main() {
	c := make(chan int)
	s := []int{7, 2, 8, -9, 4, 0}
	go sum(s[:len(s) / 2], c)
	go sum(s[len(s) / 2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}
