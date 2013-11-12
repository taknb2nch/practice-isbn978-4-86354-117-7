package main

import (
	"fmt"
	"os"
)

func main() {
	
	current, _ := os.Getwd()
	fmt.Println(current)

	os.Chdir("..")
	//os.Chdir("c:")

	current, _ = os.Getwd()
	fmt.Println(current)

}