package handlers

import (
	"net/http"

	"github.com/Vkanhan/go-rental-marketplace/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "homepage.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "aboutpage.tmpl")

}
