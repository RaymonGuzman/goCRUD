package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	Name  string
	Age   int
	Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// var person = &Person{Name: "Raymitos", Age: 26, Email: "raymon@correo.com"}

func main() {

	// http.HandleFunc()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getPeople)
	router.HandleFunc("/create", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We are on HomePage")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var db *gorm.DB
	db = connect()
	// db.Migrator().DropTable(&Person{})
	db.AutoMigrate(&Person{})

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	createdPerson := db.Create(&person)
	// db.Create(&person)
	err = createdPerson.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&person)
	}
	// fmt.Print(createdPerson)
	// json.NewEncoder(w).Encode(&person)
	// fmt.Fprint(w, "Creating user")
	// fmt.Print("Creating user")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var db *gorm.DB
	db = connect()
	// db.Migrator().DropTable(&Person{})
	db.AutoMigrate(&Person{})

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	var people []Person
	db.Find(&people)
	json.NewEncoder(w).Encode(&people)

}
