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
	c := make(chan int, 10)

	go Walk(t1, c)
	go Walk(t2, c)

	for i:= 0; i < cap(c); i++ {
		t1_v, t2_v := <-c, <-c
		fmt.Printf("t1_v: %d\n", t1_v)
		fmt.Printf("t2_v: %d\n", t2_v)
		if t1_v != t2_v {
			return false
		}
	}

	return true
}