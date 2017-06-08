// Kevin Yen-Kuan Lee
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs 是 v 的屬性, 不會改變 v 本身
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 是 v 的屬性, 且 v 是 *Vertex, 故 Scale 可以改變 v 本身
// 目前理解是這樣
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
