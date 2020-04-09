package handlers

import (
	"net/http"

	"../utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "MacServer v.1.0")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Add("x-content-type-options", "nosniff")

	utils.RenderTemplate(w, "application/index", nil)
}
