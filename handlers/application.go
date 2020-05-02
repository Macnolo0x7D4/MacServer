package handlers

import (
	"../database"
	"../utils"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "application/index", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	errorContext := make(map[string] interface{})

	source := r.URL.Query()["source"]

	if len(source[0]) > 0 {
		errorContext["Registered"] = "Your account has registered! Now, please log in to activate it."
	}

	if r.Method == "POST"{
		email := r.FormValue("email")
		password := r.FormValue("password")

		if user, err := database.Login(email, password); err != nil{
			errorContext["Error"] = err
			if errorContext["Registered"] != nil{
				errorContext["Registered"] = nil
			}
		} else{
			utils.CreateCookie(user, w)
			log.Printf("%s has logged in.", user.Username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

	utils.RenderTemplate(w, "application/login", errorContext)
}

func Logout(w http.ResponseWriter, r *http.Request){
	utils.DeleteCookie(w, r)
	http.Redirect(w,r,"/", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	errorContext := make(map[string] interface{})

	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		user := database.NewUser(username, password, email, 32, 0)

		if err := user.Valid(); err != nil{
			errorContext["Error"] = err.Error()
			log.Printf("%s wanted to register, but he/she had an issue: %s", user.Username, err.Error())
		} else{
			user.Save()
			log.Printf("%s has registered!", user.Username)
			val, _ := uuid.NewRandom()
			uid := val.String()
			http.Redirect(w, r, "/login?source=register?" + uid, http.StatusSeeOther)
		}
	}

	utils.RenderTemplate(w, "application/register", errorContext)
}