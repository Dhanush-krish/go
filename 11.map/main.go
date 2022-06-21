package main

import "fmt"

func main() {

	lookup := map[int]string{1: "one", 2: "two", 3: "three"}

	fmt.Println("elements in maps are", lookup)
	fmt.Println("1 is equivalent to:", lookup[1])
	fmt.Println("")

	//update a key
	lookup[1] = "ONE"
	fmt.Println("1 is equivalent to:", lookup[1])
	fmt.Println("")

	//iterating key, value in map
	for key, value := range lookup {
		fmt.Printf("%d - %s\n", key, value)
	}
	fmt.Println("")

	//delete key-value in map
	delete(lookup, 1)
	fmt.Println("elements in maps are", lookup)

}
