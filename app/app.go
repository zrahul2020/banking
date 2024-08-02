package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", Greet)
	mux.HandleFunc("/customers", GetAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
