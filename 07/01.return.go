package main

import (
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Open("test.txt")

	if  err != nil {

		fmt.Println("エラー:", err.Error())

		os.Exit(1)

	}

	file.Close()

	fmt.Println("ok")

}