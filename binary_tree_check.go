package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func Walk(tree *tree.Tree, c chan int) {
	for i := 0; i < cap(c); i++ {
		c <- i + 1
	}
	close(c)
}
