package main

import (
	"log"
	"net/http"
	"taqsir/handeler" 

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/order/notify", handeler.Order).Methods("POST")
	log.Println("Service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
