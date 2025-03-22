package main

import "fmt"

// 型定義はこんな感じ
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}