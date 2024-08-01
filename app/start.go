package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
