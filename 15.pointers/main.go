package main

import "fmt"

func main() {
	var value int
	var ptr *int = &value
	fmt.Println("Value inside the variable is", value)
	addpointer(&value)
	fmt.Println("Value inside the variable is ", value)
	fmt.Println("value store in ptr variable is ", *ptr)

}

func addpointer(arg *int) {
	*arg++
}
