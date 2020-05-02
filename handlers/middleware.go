package handlers

import (
	"../models"
	"../utils"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func Auth(handler handlerFunc) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuth(r){
			models.SendForbidden(w)
			return
		}

		if user := utils.GetUserByRequest(r); user != nil{
			if user.Role != 0 {
				models.SendForbidden(w)
				return
			}
		} else {
			utils.DeleteCookie(w, r)
			models.SendBadRequest(w)
			return
		}

		handler(w, r)
	})
}

func AuthHandler(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuth(r){
			models.SendForbidden(w)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

