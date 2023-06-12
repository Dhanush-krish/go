package main

import "fmt"

func main() {

	lookup := map[int]string{1: "one", 2: "two", 3: "three"}

	// creating empty map
	empty := make(map[string]int)
	fmt.Println("empty map", empty)
	fmt.Println()

	fmt.Println("elements in maps are", lookup)
	fmt.Println("1 is equivalent to:", lookup[1])
	fmt.Println("")

	//update a key
	lookup[1] = "ONE"
	fmt.Println("1 is equivalent to:", lookup[1])
	fmt.Println("")

	//go returns 0 for non existing keys

	if value, exists := lookup[5]; exists {
		fmt.Println("Value of ", 5, " is ", value)
	} else {
		fmt.Println("key ", 5, " is not in map")
	}
	fmt.Println()

	//iterating key, value in map
	for key, value := range lookup {
		fmt.Printf("%d - %s\n", key, value)
	}
	fmt.Println("")

	// iterate only keys
	for key := range lookup {
		fmt.Println(key)
	}
	fmt.Println("")

	//delete key-value in map
	delete(lookup, 1)
	fmt.Println("elements in maps are", lookup)

}
