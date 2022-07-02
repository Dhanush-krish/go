package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	fmt.Println("Go orm basics")
	dsn := "root:root@tcp(127.0.0.1:3306)/playground?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db.Open", err)
	}

	//table structure
	type Users struct {
		ID       string
		Email    string
		Password string
	}

	var user Users

	db.First(&user, "id = ?", "bd6b1f8a-41d3-404a-bf95-01dc3063e164")
	fmt.Println(user)
	fmt.Println("")

}
