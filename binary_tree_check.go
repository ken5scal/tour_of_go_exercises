package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	fmt.Println("hoge")
}

func Walk(tree *tree.Tree, c chan int) {}
