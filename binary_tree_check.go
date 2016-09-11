package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
}

func Walk(tree *tree.Tree, c chan int) {}
