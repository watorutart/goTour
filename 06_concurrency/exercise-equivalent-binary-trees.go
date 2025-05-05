package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Walk the tree in-order and send values to ch
	if t != nil {
		fmt.Println("Walked tree:", t.Value)
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// TODO: Same関数を実装
	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println("Same:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same:", Same(tree.New(1), tree.New(2)))
}
