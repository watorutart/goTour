package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// インターフェースiが具体的な型stringを保持し、基になるstringの値を変数sに代入する
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
