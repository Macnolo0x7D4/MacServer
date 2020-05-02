package utils

import (
	"../config"
	"html/template"
	"log"
	"net/http"

)

var templates = template.Must(template.New("app").ParseGlob(config.Application()))
var errorTemplate = template.Must(template.ParseFiles(config.ApplicationError()))

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.Header().Add("Server", "MacServer v.1.0")
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Add("Server", "MacServer v.1.0")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Add("x-content-type-options", "nosniff")

	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Println(err)
		RenderErrorTemplate(w, http.StatusInternalServerError)
	}
}
