package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// :=はvar宣言の代わりの代入文。暗黙的な型宣言
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}