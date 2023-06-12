package main

import "fmt"

// go has only for loop

func main() {

	// while loop implementation using for loop
	counter := 1
	for counter <= 10 {
		fmt.Println("counter using while loop -", counter)
		counter++
	}

	//c style for loop
	for iter := 1; iter <= 10; iter++ { //preincrement is not allowed in go

		if iter == 5 {
			fmt.Println("5 is skipped")
			continue
		}

		if iter == 7 {
			fmt.Println("breaking the loop at ", iter)
			break
		}

		fmt.Println("counter using for loop ", iter)
	}

	//for range
	slice := []string{"zero", "one", "two", "three", "four"}
	for idx, element := range slice {
		fmt.Println(element, "is at index", idx)
	}

}
