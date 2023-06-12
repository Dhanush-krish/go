package main

import (
	"fmt"
)

// function inside a struct
type EligibleToVote func(int) bool

// struct is used to store variables of different data types
// struct definition
type Person struct {
	Name string
	Age  int
	Mark float32
}

type Citizen struct {
	Name string
	Age  int
	Vote EligibleToVote
}

func main() {

	// have all the property values and order is same as type declaration
	person := Person{"dhanush", 24, 75.05}
	fmt.Println("values in struct", person)
	fmt.Print()

	// have few values
	person2 := Person{
		Name: "no name",
	}
	fmt.Printf("%+v", person2)
	fmt.Println("")

	//acessing property in struct
	fmt.Println("My age is", person.Age)

	citizen := Citizen{
		Name: "dummy",
		Age:  25,
		Vote: func(age int) bool {
			return age >= 18
		},
	}

	fmt.Printf("%+v", citizen.Vote(citizen.Age))

}
