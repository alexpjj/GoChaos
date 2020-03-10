package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexpjj/GoChaos/internal/pkg"
	"github.com/gorilla/mux"
)

var userRepo pkg.MongoRepository

type user struct {
	Name  string
	Email string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{email}", get).Methods("GET")
	r.HandleFunc("/user/{email}", update).Methods("PUT")
	r.HandleFunc("/user", create).Methods("POST")
	r.HandleFunc("/user/{email}", delete).Methods("DELETE")

	http.ListenAndServe(":8081", r)
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	fmt.Println(userRepo.Get("users", "email", email))
}

func update(w http.ResponseWriter, r *http.Request) {
}

func create(w http.ResponseWriter, r *http.Request) {
	var user user
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	err = userRepo.Create("users", user)
	if err != nil {
		panic(err)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
}

func init() {
	var err error
	userRepo, err = pkg.NewMongoRepository("localhost", "27017", "goChaos")

	if err != nil {
		panic(err)
	}
}
