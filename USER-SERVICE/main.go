package main

import (
	"log"
	"net/http"
	"taqsir/handeler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", handeler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handeler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handeler.DeleteUser).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
