package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 3

func main() {
	c := make(chan int)

	for i := 0; i < goroutines; i++ {

		go func(s chan<- int, x int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("処理完了:", x)

			s <- x
		}(c, i)

	}

	for i := 0; i < goroutines; i++ {
		x, _ := <- c

		fmt.Println("完了:", x)
	}

	fmt.Println("すべて完了")
}