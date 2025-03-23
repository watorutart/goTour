package main

import "fmt"

func main() {
	sum := 1
	// セミコロンを省略可能、whileの代わりにforだけを使う
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}