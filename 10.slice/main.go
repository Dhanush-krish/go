package main

import "fmt"

// collection of similar data types
// unlimited size

func main() {

	// slice initialization
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("elements in slice:", slice)

	// adding element to slice
	var slice1 []int
	slice1 := slice1.append(10)
	fmt.Println("Elements in slice1:", slice1)

}
