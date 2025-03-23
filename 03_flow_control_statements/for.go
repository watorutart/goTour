package main

import "fmt"

func main() {
	sum := 0
	// 丸括弧は不要、波括弧は常に必要
	for i :=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}