package main

import (
	"fmt"
	"os"
)

func main() {
	
	// いずれも戻り値はerrorのみ
	os.MkdirAll("a/b/c/d", 0777)
	os.MkdirAll("x/y/z", 0777)

	//os.Remove("a/b/c/d")
	os.RemoveAll("x")

	// Fileとerror
	file, err := os.Create("a/b/c/d/test")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file.Close()

	// いずれも戻り値はerrorのみ
	//os.Remove("test")
	// ディレクトリ内が空でない場合は削除出来ません
	err = os.Remove("a/b/c/d")

	if err != nil {
		fmt.Println(err)
	}
}