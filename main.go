package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portnumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("about", About)

	fmt.Println("Listening to port 8080..")
	log.Fatal(http.ListenAndServe(portnumber, nil))
}

func Home(w http.ResponseWriter, r *http.Request) {

}

func About(w http.ResponseWriter, r *http.Request) {
	
}