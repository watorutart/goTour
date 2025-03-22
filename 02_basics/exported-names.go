package main

import (
	"fmt"
	"math"
)

func main() {
	// exportされていない名前(=小文字で始まる名前)は外部のパッケージからアクセスすることは不可能
	// fmt.Println(math.pi) // error: pi is not exported by math
	fmt.Println(math.Pi)
}