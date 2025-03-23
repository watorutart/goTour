package main

import "fmt"

func main() {
	// defer: 関数のreturnするまで実行を遅延
	defer fmt.Println("world")

	fmt.Println("hello")
}