package main

import "fmt"

func main() {

	fmt.Println("Sum of two numbers =", add(10, 20))
	fmt.Println()

	//variadic function
	// You can only have one variadic parameter in a function
	//It must be the last parameter defined in the function
	fmt.Println("sum  - ", dynamic(1, 2, 3, 4))
	fmt.Println("")

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	// we can pass slice to a variadic function by exploding a slice (appending three dots at the end of string)
	fmt.Println("sum is ", dynamic(numbers...))
	fmt.Println("")

	//anonymous function
	anonymous := func(a int, b int) int {
		return a + b
	}
	fmt.Println("sum of two numbers is - ", anonymous(10, 20))
	fmt.Println("")

	//another way of calling anonymous function
	func(a int, b int) {

		fmt.Println("arguments are ", a, b)

	}(10, 20)

}

func add(arg1 int, arg2 int) int {
	return (arg1 + arg2)
}

// variadic function
func dynamic(values ...int) int {
	var result int

	for _, value := range values {
		result += value
	}
	return result
}
