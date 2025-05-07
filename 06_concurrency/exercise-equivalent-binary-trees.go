package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Walk the tree in-order and send values to ch
	var walkInOrder func(t *tree.Tree)
	walkInOrder = func(t *tree.Tree) {
		if t == nil {
			return
		}
		fmt.Println("Walked tree:", t.Value)
		walkInOrder(t.Left)
		ch <- t.Value
		walkInOrder(t.Right)
	}

	walkInOrder(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
    // 両方の木から値を受け取って比較
    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2

        // チャネルが閉じられたか確認
        if !ok1 && !ok2 {
            // 両方のチャネルが閉じられたなら、すべての値が一致したということ
            return true
        }

        // どちらか一方だけがクローズされたか、値が異なる場合
        if !ok1 || !ok2 || v1 != v2 {
            return false
        }
    }
}

func main() {
   // Walkのテスト
   ch := make(chan int)
   go Walk(tree.New(1), ch)

   // チャネルから値を受け取って表示
   fmt.Print("Tree values: ")
   for v := range ch {
	   fmt.Print(v, " ")
   }
   fmt.Println()

   // Sameのテスト
   fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))
   fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
}
