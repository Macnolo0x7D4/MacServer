package utils

import (
	"../database"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

const (
	cookieAuthName = "session"
)

var Sessions = struct {
	m map[string]*database.User
	sync.RWMutex
}{m: make(map[string]*database.User)}

func CreateCookie(user *database.User, w http.ResponseWriter){
	Sessions.Lock()
	defer Sessions.Unlock()

	num, _ := uuid.NewRandom()
	value := num.String()

	Sessions.m[value] = user
	cookie := &http.Cookie{
		Name: cookieAuthName,
		Value: value,
		Path: "/",
	}

	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	Sessions.Lock()
	defer Sessions.Unlock()

	delete(Sessions.m,GetCookieUUID(r))
	cookie := &http.Cookie{
		Name:   cookieAuthName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func GetCookieUUID(r *http.Request) string{
	cookie, err := r.Cookie(cookieAuthName)

	if err != nil{
		return ""
	}

	return cookie.Value
}

func IsAuth(r *http.Request) bool{
	return GetCookieUUID(r) != ""
}

func GetUserByRequest(r *http.Request) *database.User{
	Sessions.Lock()
	defer Sessions.Unlock()

	value := GetCookieUUID(r)

	if user, ok := Sessions.m[value]; ok{
		return user
	} else{
		return nil
	}
}