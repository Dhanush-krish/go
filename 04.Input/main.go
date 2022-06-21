package main

import (
	"fmt"
)

func main() {

	// we can use scan to get multiple space seperated input
	var name, age string
	fmt.Println("Enter your name and age")
	fmt.Scan(&name, &age)
	fmt.Println("your inputs are", name, age)

	// we can use scanln to get line input
	var name1 string
	fmt.Println("Enter your name and age")
	fmt.Scanln(&name1)
	fmt.Println("your inputs are", name1)

	//c style input
	var name3 string
	var age3 int
	fmt.Println("Enter name and age")
	fmt.Scanf("%s %d", &name3, &age3)
	fmt.Printf("your name is %v of type %T and age %v of type %T \n", name3, name3, age3, age3)

}
