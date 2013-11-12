package main

import "fmt"

func main() {
	
	capitals := make(map[string]string)

	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンD.C."
	capitals["イギリス"] = "ロンドン"
	capitals["フランス"] = "パリ"

	fmt.Println(capitals["日本"])
	fmt.Println(capitals["アメリカ"])
	fmt.Println(capitals["イギリス"])
	fmt.Println(capitals["フランス"])
	// 未登録の場合はゼロ値が返ります。
	// 文字列型の場合はnilではなく""空文字列
	fmt.Println(capitals["カナダ"])

	// 順序は保証されません
	for country, capital := range capitals {
		fmt.Println(country, capital)
	}

	// 登録確認するだけなら
	// _, ok := capitals["中国"]
	capital, ok := capitals["中国"]

	if ok {
		fmt.Println("登録あり:", capital)
	} else {
		fmt.Println("未登録:", capital)
	}

	delete(capitals, "日本")

	for country, capital := range capitals {
		fmt.Println(country, capital)
	}

	// 存在しないキーを削除した場合はエラーにはなりません。
	delete(capitals, "日本")	
}