package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// LIFOの順番で実行される
		defer fmt.Println(i)
	}

	fmt.Println("done")
}