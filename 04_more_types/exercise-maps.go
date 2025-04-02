package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

// FIXME: ゴリ押しで行けたんで後で見返したい
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var tmp string = ""
	fmt.Println(tmp)
	for _, c := range s {
		if c == ' ' {
			m[tmp]++
			tmp = ""
			continue
		} else {
			tmp = tmp + string(c)
		}
	}
	m[tmp]++
	return m
}

func main() {
	wc.Test(WordCount)
}