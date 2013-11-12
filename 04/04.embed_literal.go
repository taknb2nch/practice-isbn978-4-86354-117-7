package main

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
	Person
}

func main() {
	
	e1 := Employee{1, Person{"Jack", 23}}


	// 構造体リテラルを用いて初期化する場合、
	// 埋め込みフィールドは直接初期化できない
	//NG: 
	// e2 := Employee{id:2, name:"hogehoge", age:30}

	// OK:
	var e2 Employee
	e2.id = 10
	e2.name = "Emi"
	e2.age = 18

	fmt.Println(e1, e2)

}