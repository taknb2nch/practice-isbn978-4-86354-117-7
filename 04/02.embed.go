package main

import "fmt"

type embedded struct {
	i int
	_ byte
	string
}

func (x embedded) doSomething() {
	fmt.Println("test.doSomething")
}

type test struct {
	embedded
}

func main() {

	var x test

	x.doSomething()
	
	x.string = "あいうえお"

	fmt.Println(x.i, x.string)

	fmt.Println(x)
}