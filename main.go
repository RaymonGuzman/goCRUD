package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	Name  string
	Age   int
	Email string `gorm:"type:varchar(100);unique_index"`
}

var person = &Person{Name: "Raymitos", Age: 26, Email: "raymon@correo.com"}

func main() {
	var db *gorm.DB
	db = connect()
	db.Migrator().DropTable(&Person{})
	db.AutoMigrate(&Person{})
	db.Create(person)
	fmt.Println("Funcionooo")

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()
}
