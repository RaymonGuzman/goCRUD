package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	// handleRequest()
// 	testing.Testing()
// 	Pruebita()
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	fmt.Println("Probando el server")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
