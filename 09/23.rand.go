package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())

	fmt.Println(rand.Int63n(3))

	fmt.Println(rand.Float32())

	fmt.Println(rand.Float64())	

}