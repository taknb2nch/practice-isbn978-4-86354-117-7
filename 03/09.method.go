package main

import "fmt"

type myType int

func main() {
	var v myType = 2345

	v.println()

	fmt.Println(v.square())
}

func (value myType) println() {
	fmt.Println(value)
}

func (value myType) square() myType {
	return value * value
}