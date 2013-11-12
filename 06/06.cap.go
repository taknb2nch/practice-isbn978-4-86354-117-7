package main

import "fmt"

func main() {
	
	values := [...]int { 0, 1, 2, 3, 4 }

	s1 := values[1:4]

	fmt.Println(s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	s2 := s1[1:4]

	fmt.Println(s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))

    // NG: もとの配列の最大インディックスを超えるため
	//s3 := s2[1:4]
	s3 := s2[1:3]

	fmt.Println(s3)
	fmt.Println("len:", len(s3))
	fmt.Println("cap:", cap(s3))
}