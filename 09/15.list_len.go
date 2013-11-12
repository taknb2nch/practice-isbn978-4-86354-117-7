package main

import (
	"container/list"
	"fmt"
)

func main() {
	
	l := list.New()

	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println(l.Len(), l)
}