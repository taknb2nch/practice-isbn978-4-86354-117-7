package main

import "fmt"

func main() {
	
	var i interface{} = "test"

	var s = i.(string)

	fmt.Println(s)

}