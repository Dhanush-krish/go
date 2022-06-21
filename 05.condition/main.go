package main

import "fmt"

func main() {
	var score int

	fmt.Println("Enter your mark ")
	fmt.Scanf("%d", &score)

	if score > 80 {
		fmt.Println("Distinction")
	} else if score > 60 {
		fmt.Println("Average")
	} else {
		fmt.Println("Below Average")
	}

	//with declaration statements         **relational operators is must in evaluation**
	if num := 2; num%2 == 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}
