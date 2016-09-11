package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t1))
}

func Walk(tree *tree.Tree, ch chan int) {
	_walk(tree, ch)
	close(ch)
}

func _walk(tree *tree.Tree, ch chan int) {
	if (tree.Left != nil) {
		_walk(tree.Left, ch)
	}
	if (tree.Right != nil) {
		_walk(tree.Right, ch)
	}

	ch <- tree.Value
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i:= range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}