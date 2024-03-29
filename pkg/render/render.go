package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./template/"+tmpl, "./template/base.layout.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template", err)
		return

	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	//check to see if we already have template in our cache
	_, inMap := tc[t]
	if !inMap {
		//create the template
		log.Println("creating template and adding it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("error creating template", err)
		}
	} else {
		//we have template in the cache
		log.Println("using the cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error executing template:", err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("/template/%s", t),
		"./template/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//add template to cache (map)
	tc[t] = tmpl

	return nil

}
