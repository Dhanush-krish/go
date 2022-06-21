package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Enter two space seperated integers ")
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("Sum of two numbers =", add(num1, num2))

}

func add(arg1 int, arg2 int) int {
	return (arg1 + arg2)
}
