package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 返り値を省略することもできる。ただし、可読性に難ありのため長い関数では非推奨
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split(18))
}