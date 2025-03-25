package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 配列のスライス(可変長)
	// aの要素のうち1-3番目を含むスライスを作成
	var s []int = primes[1:4]
	fmt.Println(s)
}