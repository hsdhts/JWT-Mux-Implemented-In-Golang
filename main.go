package main

import (
	"github.com/gorilla/mux"
	"github.com/jeypc/go-jwt-mux/controllers/authcontroller"
	"github.com/jeypc/go-jwt-mux/models"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
