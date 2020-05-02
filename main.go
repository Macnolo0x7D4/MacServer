package main

import (
	"./config"
	"./database"
	"./handlers"
	v1 "./handlers/api/v1"
	"bufio"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var Server *http.Server

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.Logout)
	r.HandleFunc("/register", handlers.Register).Methods("GET", "POST")

	r.Handle("/api/v1/users/", handlers.Auth(v1.GetUser)).Methods("GET")
	r.Handle("/api/v1/users/", handlers.Auth(v1.CreateUser)).Methods("POST")
	r.Handle("/api/v1/users/{id:[0-9]+}", handlers.Auth(v1.GetUser)).Methods("GET")
	r.Handle("/api/v1/users/{id:[0-9]+}", handlers.Auth(v1.UpdateUser)).Methods("PUT")
	r.Handle("/api/v1/users/{id:[0-9]+}", handlers.Auth(v1.DeleteUser)).Methods("DELETE")

	assets := http.FileServer(http.Dir(config.Assets()))
	statics := http.StripPrefix("/assets", assets)

	r.PathPrefix("/assets/").Handler(statics)

	database.CreateConnection()

	//database.CreateTable()

	Server = &http.Server{
		Addr: config.GetUrlHttp(),
		Handler: r,
	}

	log.Println("Listening:", "http://" + config.GetUrlHttp())

	go cli()

	log.Fatal(Server.ListenAndServe())
}

func cli(){
	reader := bufio.NewReader(os.Stdin)

	for{
		cmd, _, _ := reader.ReadLine()

		if string(cmd) == "stop" {

			log.Println("Closing MySQL Communication...")
			database.CloseConnection()

			log.Println("Stopping HTTP server...")

			Server.Shutdown(context.Background())

			log.Println("Done! Bye.")

			os.Exit(0)
		}
	}
}
