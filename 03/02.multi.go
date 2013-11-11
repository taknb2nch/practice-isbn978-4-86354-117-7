package main

import "fmt"

func main() {
	add, sub, mul, div := calc(1, 2)

	fmt.Println(add, sub, mul, div)
}

// 複数戻り値の場合は括弧で囲む必要あり
func calc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}