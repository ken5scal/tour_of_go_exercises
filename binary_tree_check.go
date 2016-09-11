package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t1))
}

func Walk(tree *tree.Tree, c chan int) {
	_walk(tree, c)
	close(c)
}

func _walk(tree *tree.Tree, c chan int) {
	if (tree.Left != nil) {
		_walk(tree.Left, c)
	}
	if (tree.Right != nil) {
		_walk(tree.Right, c)
	}

	c <- tree.Value
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i:= range c1 {
		if i != <-c2 {
			return false
		}
	}

	return true
}