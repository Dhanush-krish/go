package main

import "fmt"

type Profile struct {
	Name string
	age  int
}

func main() {
	var value int
	var ptr *int = &value

	fmt.Println("Value inside the variable is", value)

	addpointer(&value)

	fmt.Println("Value inside the variable is ", value)

	fmt.Println("value store in ptr variable is ", *ptr)

	//struct pointer
	person := Profile{"sherlock", 45}
	var sptr = &person
	fmt.Println(sptr.Name, sptr.age)

}

func addpointer(arg *int) {
	*arg++
}
