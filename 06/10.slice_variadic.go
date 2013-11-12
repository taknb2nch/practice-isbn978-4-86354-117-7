package main

import "fmt"

func main() {
	
	s := []string { "a", "b", "c" }

	test2(s)

	test(s...)

	test("a", "b", "c")
}

func test(s ...string) {
	fmt.Println(len(s), s)
}

func test2(s ...[]string) {
	fmt.Println(len(s), s)
}