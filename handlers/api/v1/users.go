package v1

import (
	"../../../database"
	"../../../models"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, database.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := getVars(r)

	if user := database.GetUserById(userId); user.Id == 0 {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &database.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if _, err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	userResponse := database.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.Username = userResponse.Username
	user.Password = userResponse.Password
	user.Email = userResponse.Email

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	} else {
		user.Save()
		models.SendData(w, user)
	}
}

func getUserByRequest(r *http.Request) (*database.User, error) {
	userId := getVars(r)

	user := database.GetUserById(userId)

	if user.Id == 0 {
		return user, errors.New("User not exists")
	} else {
		return user, nil
	}
}

func getVars(r *http.Request) int {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	return userId
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)

	if err != nil {
		models.SendNotFound(w)
	} else {
		user.Delete()
		models.SendNoContent(w)
	}
}