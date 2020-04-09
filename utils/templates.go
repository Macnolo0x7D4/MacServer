package utils

import (
	"html/template"
	"log"
	"net/http"

	"../config"
)

var templates = template.Must(template.New("app").ParseGlob(config.Application()))
var errorTemplate = template.Must(template.ParseFiles(config.ApplicationError()))

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Println(err)
		RenderErrorTemplate(w, http.StatusInternalServerError)
	}
}
