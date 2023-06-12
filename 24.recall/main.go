package main

import "fmt"

func main() {

	//print output
	fmt.Println("I'm learning go lang")
	fmt.Println("")

	//get input from console
	var age int
	fmt.Scanf("%d", &age)
	fmt.Printf("My age is %d\n", age)

	//condition
	if age < 18 {
		fmt.Println("you are not eligible to vote")
	} else {
		fmt.Println("You are eligible to vote")
	}

	//switch statement
	switch {
	case age < 10:
		fmt.Println("your age is below 10")
	case age > 50:
		fmt.Println("your age is above 50")
		fallthrough
	default:
		fmt.Println("I'm in default statement")
	}

	//looping
	for iter := 1; iter <= 10; iter++ {
		fmt.Println(iter)
	}

	//array
	var array [10]int
	array[0] = 10
	fmt.Println(array)

	var str string = "checking string indexing"

	fmt.Printf("character at postition 0 is %c", str[0])

}
