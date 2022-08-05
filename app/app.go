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
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//define routes
	router.HandleFunc("/cutomers", ch.getCustomers).Methods(http.MethodGet)
	//starting server
	http.ListenAndServe("localhost:5005", router)
}
