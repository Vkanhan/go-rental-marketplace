package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vkanhan/go-rental-marketplace/pkg/config"
	"github.com/Vkanhan/go-rental-marketplace/pkg/handlers"
	"github.com/Vkanhan/go-rental-marketplace/pkg/render"
)

const portnumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Application starting on port %s..", portnumber)
	// log.Fatal(http.ListenAndServe(portnumber, nil))

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
