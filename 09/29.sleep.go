package main

import (
	"fmt"
	"time"
)

func main() {
	
	fmt.Println(time.Now())

	time.Sleep(time.Second * 5)

	fmt.Println(time.Now())

	time.Sleep(time.Minute * 1)

	fmt.Println(time.Now())

}