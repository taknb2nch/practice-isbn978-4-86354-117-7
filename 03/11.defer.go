package main

import "fmt"

func main() {
	
	fmt.Println("開始")

	defer delay()

	defer fmt.Println("delay2")

	defer func(x int) {
		fmt.Println("delay3", x)
	}(1)

	fmt.Println("終了")

}

func delay() {
	fmt.Println("delay1")
}