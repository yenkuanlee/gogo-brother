// Kevin Yen-Kuan Lee

package main

import (
	"fmt"
	"math"
)

type MyFloat float64
// 一定要先定義 type, 才能使用他的屬性

func (f MyFloat) Abs() float64 { // Abs() 這個 function 是 MyFloat 的屬性
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2) // f 的 type 為 MyFloat
	fmt.Println(f.Abs()) // 呼叫 f 的屬性
}
