package main

import "fmt"

//struct => similar to class
type Student struct {
	Name string
	Age  int
}

// method
func (s Student) greetStudent(message string) {
	fmt.Printf("%v %v\n", message, s.Name)
}

func main() {

	obj := Student{"Dhanush", 24}
	fmt.Printf("%+v\n", obj)

	obj.greetStudent("hello")

}
