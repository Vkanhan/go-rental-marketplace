package main

import (
	"net/http"

	"github.com/Vkanhan/go-rental-marketplace/pkg/config"
	"github.com/Vkanhan/go-rental-marketplace/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
