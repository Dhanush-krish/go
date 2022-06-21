package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter a number between 1-10")
	fmt.Scanf("%d", &number)

	switch number { //declaration is possible as in condition statement
	case 1:
		fmt.Println("ONE") //to execute the next case statement
		fallthrough
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	case 4:
		fmt.Println("FOUR")
	case 5:
		fmt.Println("FIVE")
	case 6, 7, 8, 9, 10: //multiple expression check
		fmt.Println("Second half")
	default:
		fmt.Println("Number not in range 1-10 :", number)
	}
}
