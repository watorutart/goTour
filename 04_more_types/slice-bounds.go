package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	// スライスの上限・下限を省略可能
	s = s[:2]
	fmt.Println(s)

	// 配列１番目から最後まで
	s = s[1:]
	fmt.Println(s)
}