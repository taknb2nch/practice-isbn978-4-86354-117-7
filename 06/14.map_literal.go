package main

import "fmt"

func main() {
	
	// 最後の行に閉じ括弧を付けるか、最後の行もコンマを付けるかのいずれか
	capitals := map[string]string {
		"日本" : "東京",
		"アメリカ" : "ワシントンD.C.",
		"イギリス" : "ロンドン",
	}

	fmt.Println(capitals)

}