package main

import "fmt"

func fibonacci() func() int {
	// n1, n2がstaticにfunc内で保持されて更新されていってるイメージ？
	n1, n2 := 0, 1
	return func() int {
		n := n1
		n1, n2 = n2, n1+n2
		fmt.Println("中間データ(n/n1/n2): ", n, n1, n2)
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
}
