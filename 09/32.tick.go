package main

import (
	"fmt"
	"time"
)

func main() {
	
	// no need shut down the ticker
	ch := time.Tick(time.Second * 2)

	for i := 0; i < 10; i++ {
		t := <-ch

		fmt.Println(t)
	}	

}