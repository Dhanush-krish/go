package main

import (
	"fmt"

	sql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := sql.Config{
		User:                 "developer",
		Passwd:               "password",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "development",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: config.FormatDSN()}))
	if err != nil {
		fmt.Println("db.Open", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		//TODO need to add logging
		fmt.Println(db, err)
	}

	if err := sqlDB.Ping(); err != nil {
		//TODO need to add logging
		fmt.Println(db, err)
	}

	// tables models
	type Order struct {
		gorm.Model
	}

	type Card struct {
		gorm.Model
		Name string
	}

	type User struct {
		gorm.Model
		Name   string
		CardId int
		Card   Card
	}

	db.AutoMigrate(&User{})

}
