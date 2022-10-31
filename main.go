package main

import "fmt"

func main() {
	fmt.Println("hello GO")
	/*
	   understanding variables and namings

	   # variables can be declared using var keyword followed by variable name and the type of varaible

	   var num1 int = 10
	   var name string = "kushal"
	   var num2 float64=9.55
	   var state bool = false

	*/

	/*
		The implicit declaration of variable type

		name := "kushal"
		number :=5
		number2 :=8.8

		this is useful, since we can define variable without explicitly defining their datatypes
	*/
	/*
	   the printf and scanf

	   val := 5
	   name:="kushal"
	   fmt.Printf("Name is %v ",name)
	   fmt.Printf("the value is %v ",val)

	   --> the %v refers to the value of the variable

	   #taking user input in go

	   fmt.Scan()

	   for ex: var name string

	   fmt.Scan(&name)  -> taking user input and storing in the memory location


	*/
	var state bool = false
	fmt.Println(state)

	name := "kushal"
	surname := "subedi"
	fmt.Println("Full Name is", name+"", surname)

	val1 := 2
	val2 := 5
	fmt.Printf("%v \n", val1+val2)
}