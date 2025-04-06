package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
// ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。
// 一般的には、値レシーバ、または、ポインタレシーバのどちらかですべてのメソッドを与え、混在させるべきではありません。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling:%+v, Abs: %v\n",v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling:%+v, Abs: %v\n",v, v.Abs())
}