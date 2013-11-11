package main

import "fmt"

func init() {
	fmt.Println("初期化処理1")
}

func init() {
	fmt.Println("初期化処理2")
}

func main() {
	fmt.Println("メイン処理")
}