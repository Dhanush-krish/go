package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//struct is used to store variables of different data types

	// struct definition
	type Person struct {
		Name string
		Age  int
		Mark float32
	}

	person := Person{"dhanush", 24, 75.05}
	fmt.Println("values in struct", person)
	fmt.Print()

	//acessing property in struct
	fmt.Println("My age is", person.Age)

	// nested struct
	//json => regions:{{"us":{"ref_user_id":82631 }}}

	type RefId struct {
		RefUserId int `json:"ref_user_id"`
	}

	type Region struct {
		Region RefId `json:"us"`
	}

	data := []byte(`{"us":{"ref_user_id":82631 }}`)

	fmt.Println(string(data))

	var regions Region
	err := json.Unmarshal(data, &regions)
	fmt.Println(err)

	fmt.Printf("%+v", regions)
}
