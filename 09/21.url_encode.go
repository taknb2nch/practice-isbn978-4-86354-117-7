package main

import (
	"fmt"
	"net/url"
)

func main() {
	
	data := "abcあいう/:"

	enc := url.QueryEscape(data)

	fmt.Println(enc)

	// string, error
	dec, _ := url.QueryUnescape(enc)

	fmt.Println(dec)

}