package main

import "fmt"

func main() {
	//struct is used to store variables of different data types

	// struct definition
	type Person struct {
		Name string
		Age  int
		Mark float32
	}

	person := Person{"dhanush", 24, 75.05}
	fmt.Println("values in struct", person)
	fmt.Print()

	//acessing property in struct
	fmt.Println("My age is", person.Age)
}
