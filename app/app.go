package app

import (
	"net/http"

	"github.com/saurabhsingh1408/banking_app/domain"
	"github.com/saurabhsingh1408/banking_app/service"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//define routes
	router.HandleFunc("/customers", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//starting server
	http.ListenAndServe("localhost:5005", router)
}
