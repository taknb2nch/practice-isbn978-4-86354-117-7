package main

import "fmt"

func main() {
	
	values := [...]int { 0, 1, 2, 3, 4 }

	double2(values)

	fmt.Println(values)

	double(values[:])

	fmt.Println(values)


	vals := [...]int { 1, 2, 3, 4, 5 }

	vals2 := vals[:]

	vals2[0] = 100

	fmt.Println(vals)
	fmt.Println(vals2)

}

func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}

// 配列型は長さも含めて1つの値型になる
func double2(values [5]int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}