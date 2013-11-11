package main

import "fmt"

type myType int

func main() {

	var x myType = 0

	x.setByValue(1)
	fmt.Println(x)

	x.setByPointer(5)
	fmt.Println(x)	

}

func (value myType) setByValue(newValue myType) {
	value = newValue
}

func (value *myType) setByPointer(newValue myType) {
	*value = newValue
}
