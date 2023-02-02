package main

import (
	// "log"

	"github.com/seyiadel/otp-go/otp"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func main(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Connect any database for Postgresql, Sqlite
	if err != nil{
		panic("failed to connect to database")
	}

	otp.Generate(db,"qwerty@gmail.com",4) //Parameters: database, email, otp length

	}