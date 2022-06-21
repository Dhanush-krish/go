package main

import "fmt"

func main() {

	//string
	var name string = "dhanush"
	fmt.Println("My name is : ", name)

	//integer
	var age int = 24
	fmt.Println("My age is : ", age)

	//boolean
	var flag bool = false
	fmt.Println("Married : ", flag)

	//variable
	var char = "some times"
	fmt.Println("Good boy ? ", char)

	//constant
	const CHAR = "Bad boy"
	fmt.Println("Type : ", CHAR)

	// default value
	var integer int
	fmt.Println("default value for integer : ", integer)

	// default value for string
	var str string
	fmt.Println("default value for string : ", str)

	// local variable
	local := "yes"
	fmt.Println("are you a local variable : ", local)

}
