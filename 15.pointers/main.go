package main

import "fmt"

func main() {
	var value int
	// var ptr *int = &value
	fmt.Println("Value inside the variable is", value)
	addpointer(&value)
	fmt.Println("Value inside the variable is ", value)

}

func addpointer(arg *int) {
	*arg++
}
