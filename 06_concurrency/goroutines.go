package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i<5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// goroutineを使って並行処理
	go say("world")
	// 通常の関数呼び出し
	say("hello")
}
