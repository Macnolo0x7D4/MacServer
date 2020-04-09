package main

import (
	"log"
	"net/http"

	v1 "./handlers/api/v1"

	"./config"
	"./database"
	"./handlers"
	"github.com/gorilla/mux"
)

/*func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request){
	r.URL.Query()

	//uri := CreateURL(r.URL.Query.Get("href"))
	uri := "LOLOLO"

	if len(uri) != 0 {
		fmt.Fprintf(w, "Hello " + this.name + ", you will be redirected to: " + uri)
	}else {
		fmt.Fprintf(w, "Hello " + this.name)
	}
}*/

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/", v1.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")

	assets := http.FileServer(http.Dir("../www/application/assets"))
	statics := http.StripPrefix("/assets", assets)

	r.PathPrefix("/assets/").Handler(statics)

	database.CreateConnection()

	/*server := &http.Server{
		Addr:    config.GetUrlWebserver(),
		Handler: r,
	}*/

	log.Println("Listening:", config.ServerPort())

	//log.Fatal(server.ListenAndServe)
	log.Fatal(http.ListenAndServe(":8000", r))
	database.CloseConnection()
}
