package main

import "fmt"

func main() {

	type Profile struct {
		name string
	}

	fmt.Printf("%+v", Profile{"dhanush"})

}
