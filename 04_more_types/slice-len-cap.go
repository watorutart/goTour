package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// capを超えるとエラー
	// s = s[:7]
	// printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// len: スライスの長さ, cap: スライスの容量
	// スライスの長さは要素数, 容量はスライスの最初の要素から数えた要素数
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
