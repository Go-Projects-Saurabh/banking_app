package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/cutomers", getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/cutomers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/cutomers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) //URL segments
	//starting server
	http.ListenAndServe("localhost:5005", router)
}
