package main

import "fmt"

func main() {
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念日")
	holiday(3, "春分の日")
}

func holiday(month int, days ...string) {
	
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))

	// 複数の戻り値を受け取る場合後の値は省略できるが、前の値は省略できない
	// NG: , day := range days
	// OK: x := range days
	// 
	// 配列をrangeで取得した場合は インディックス, 値 となる
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println()
}