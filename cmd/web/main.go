package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vkanhan/go-rental-marketplace/pkg/handlers"
)

const portnumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Listening to port 8080..")
	log.Fatal(http.ListenAndServe(portnumber, nil))
}
