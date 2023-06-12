package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Strings in GO"
	fmt.Printf("accessing first element in str: %c\n", str[0])

	// finding the length of the string
	fmt.Println("length of the string is", len(str))

	// converting string to lowercase
	fmt.Println("converting to lower case", strings.ToLower(str))

	// converting string to uppercase
	fmt.Println("converting to lower case", strings.ToUpper(str))

	// converting string to array
	fmt.Println("converting to lower case", strings.Split(str, " "))

	// string replace
	fmt.Println("converting to lower case", strings.Replace(str, "GO", "went", 1))

}
