package main

import "fmt"

func main() {
	add, sub, mul, div := calc(1, 2)

	fmt.Println(add, sub, mul, div)
}

func calc(a int, b int) (add int, sub int, mul int, div float32) {

	// 名前指定の戻り値の場合は　:= としてはいけません。
	// NG: add := a + b
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	// 名前指定の戻り値の場合も return は必要です。
	return

	// 普通に return に戻り値を書いても大丈夫です。
	// 戻り値変数に値を代入するのと、return文の両方を書いた場合は、returnが優先されます。
	//return a + b, a - b, a * b, float32(a) / float32(b)
}