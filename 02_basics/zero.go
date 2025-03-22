package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// 変数の初期値はゼロ値
	// ゼロ値は型によって異なる
	// int: 0
	// float64: 0.0
	// bool: false
	// string: ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}