package main

import "fmt"

// collection of similar data types
// fixed size

// array ddefinition template
// var array_variable = [size]datatype{elements of array}

func main() {

	//array initialization and declaration
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("content in int array :", arr)

	// array initialization
	var arr1 [2]string
	arr1[0] = "dhanush"
	arr1[1] = "age"
	fmt.Println("content in string array :", arr1)

	//printing with quots
	fmt.Printf("with quotes - %q \n", arr1)

	//zero based index
	fmt.Println("first element in the array", arr1[0])
	fmt.Println("")

	//acessing all the element in array
	for index := 0; index < len(arr); index++ {
		fmt.Printf("%d is at position %d normal array\n", arr[index], index)
	}
	fmt.Println("")

	//acessing elements using for range ==> 0 based indexing
	for index, value := range arr {
		fmt.Printf("%d is at position %d using range\n", value, index)
	}
}
