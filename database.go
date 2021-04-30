package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Person struct {
// 	gorm.Model

// 	Name  string
// 	Age   int
// 	Email string `gorm:"type:varchar(100);unique_index"`
// }

var db *gorm.DB
var err error

// var person = &Person{Name: "Raymon", Age: 26, Email: "raymon@correo.com"}

// Pruebita ...
func connect() *gorm.DB {
	godotenv.Load()
	// testing.Testing()
	// fmt.Println("somthing")
	// // os.Exit(0)
	// Pruebita()
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	// fmt.Println(user, dbPort, host, dbName, password)
	// os.Exit(3)

	//Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", host, user, dbName, password, dbPort)

	// Openning connection to database
	// db, err = gorm.Open(dialect, dbURI)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Close connection to database when the main function finishes

	// db.AutoMigrate(&Person{})
	// db.Create(person)

	// sqlDB, err := db.DB()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer sqlDB.Close()
	return db
}
