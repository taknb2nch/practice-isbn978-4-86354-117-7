package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	
	file, err := os.OpenFile("test.csv", os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	w := csv.NewWriter(file)

	w.Write([]string { "No", "商品", "価格", "備考" })
	w.Write([]string { "1", "キャベツ", "150", "とれたて朝取れ,新鮮" })
	w.Write([]string { "2", "にんじん", "100", "" })
	w.Write([]string { "3", "サンマ", "99", "今日の\"特価品\"です" })
	w.Write([]string { "4", "まぐろ", "1025", "青森県\n大間産" })

	w.Flush()

}