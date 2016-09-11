package main

import "fmt"
import "golang.org/x/tour/tree"

func main() {

	//for i := 0; i < cap(c); i++ {
	//	fmt.Println(<-c)
	//}
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Println(t1.String())
	fmt.Println(t2.String())

	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t1))
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
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i:= 0; i < cap(c1); i++ {
		t1_v, t2_v := <-c1, <-c2
		if t1_v != t2_v {
			return false
		}
	}

	return true
}