package userapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{email}", get).Methods("GET")
	r.HandleFunc("/user/{email}", update).Methods("PUT")
	r.HandleFunc("/user/{email}", add).Methods("POST")
	r.HandleFunc("/user/{email}", delete).Methods("DELETE")

	http.ListenAndServe(":80", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}

func add(w http.ResponseWriter, r *http.Request) {
}

func delete(w http.ResponseWriter, r *http.Request) {
}

func init() {
}
