package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {
	c := make(chan int, 10)
	go Walk(tree.New(1), c)

	//for i := 0; i < cap(c); i++ {
	//	fmt.Println(<-c)
	//}
	result := Same(tree.New(1), tree.New(2))
	fmt.Println(result)
}

func Walk(tree *tree.Tree, c chan int) {
	if (tree.Left != nil) {
		Walk(tree.Left, c)
	}
	if (tree.Right != nil) {
		Walk(tree.Right, c)
	}

	c <- tree.Value
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	return true
}