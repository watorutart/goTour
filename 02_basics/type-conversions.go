package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// float64型に変換して計算しないとエラー
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, f, z)
}