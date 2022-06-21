package main

import "fmt"

// collection of similar data types
// unlimited size

func main() {

	// slice initialization
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("elements in slice:", slice[0:3])

	// adding element to slice
	var slice1 []int
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20)
	fmt.Println("Elements in slice1:", slice1)
	fmt.Println("")

	for _, index := range slice {
		fmt.Printf("value at index %d is %d\n", index, slice[index-1])
	}

}
