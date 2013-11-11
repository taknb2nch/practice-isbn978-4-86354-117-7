package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {

		for i := 0; i < 10; i++ {

			select {

			case ch1 <- i :
				fmt.Println("ch1への送信完了")

			case ch2 <- "test" + strconv.FormatInt(int64(i), 10):
				fmt.Println("ch2への送信完了")
			}

		}

		os.Exit(0)
	}()
	
	for {

		select {

		case val := <-ch1:
			fmt.Println("cｈ1からの受信完了:", val)

		case text := <-ch2:
			fmt.Println("ch2からの受信完了:", text)

		}

	}
}