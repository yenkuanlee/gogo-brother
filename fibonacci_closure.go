// Kevin Yen-Kuan Lee
// Do by myself
package main

import "fmt"

// 廢氏數列
func fibonacci() func() int {
	x := 0
	fib := 0
	return func() int{
		if fib==0{ // 初始值為 0
			fib = 1
			return 0
		}else{
			tmp := fib+x // 最新的 ： 前一個加前二個
			x = fib // 更新前一個
			fib = tmp // 更新最新
			return fib 
		}
	}
	// 重點在於宣告 f := fibonacci() 後 x 與 fib 會一直累計
}

func main() {
	f := fibonacci()
	// f 宣告後, 裡面的 x 與 fib 會一直變動
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
