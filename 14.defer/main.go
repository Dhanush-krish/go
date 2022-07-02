package main

import "fmt"

// defer statements will execute at the end of all the executions
// First differ will execute at the end

func main() {
	defer fmt.Println("Hello")
	fmt.Println("World")

	defer myCounter()
}

func myCounter() {
	for counter := 0; counter < 5; counter++ {
		defer fmt.Println(counter)

	}
}
