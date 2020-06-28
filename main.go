package main

import (
	"github.com/jinzhu/gorm"

_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type User struct {
	gorm.Model
	firstName string
	lastName string
}



func main()  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to the database")
	}
	defer db.Close()



	//Migrate the schema
	db.AutoMigrate(&User{})

	//Create
	db.Create(&User{firstName: "Jane", lastName: "Doe"})

	//Read
	var user User
	db.First(&user, 1)
	db.First(&user, "firstName = ?", "Jane")

	//Update - update User's first name to John
	db.Model(&user).Update("firstName", "John")

	//Delete the user
	db.Delete(&user)
}
