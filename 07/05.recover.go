package main

import "fmt"

func main() {
	
	f1(false)
	f1(true)

}

func f1(p bool) {
	
	defer func() {

		fmt.Println("start defer function")

		if err := recover(); err != nil {
			fmt.Println("Stop at Panic!:", err)
		}

		fmt.Println("finish defer function")

	}()

	if p {
		panic("Panic Ocuured!")
	}

}