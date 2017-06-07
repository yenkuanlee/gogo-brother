package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// 如果有人宣告 a := adder(), a 是一個 function (因為 adder 的回傳型態為 func)
	// a 裡面的 sum 初始值為 0, 之後就會一直變動, 像是 Java 裡面的 static
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	// pos 中的 sum 初始值為 0
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // 每次呼叫 pos, 他的 sum 就會累計下去
			neg(-2*i),
		)
	}
}
