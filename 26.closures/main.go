package main

import "fmt"

// closure is a nested function that allows access to outer function variable
// even after the outer function is closed

func main() {

	even := add()
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
}

func add() func() int {
	num := 0

	//closure function
	return func() int {
		num += 2
		return num
	}
}
