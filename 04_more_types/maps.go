package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// マップのゼロ値は nil です。 nil マップはキーを持っておらず、またキーを追加することもできません。
	// make 関数は指定された型のマップを初期化して、使用可能な状態で返します。
	m = make(map[string]Vertex)
	m["Bell lab"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell lab"])
}