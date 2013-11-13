package main

import (
	"fmt"
	"time"
)

func main() {
	
	loc, _ := time.LoadLocation("Local")

	time := time.Date(2003, 2, 28, 14, 53, 26, 231, loc)
	//time := time.Date(2003, 2, 29, 14, 53, 26, 231, loc)
	// Output: 
	// 2003-03-01 14:53:26.000000231 +0900 +0900

	fmt.Println(time)

}