package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/saurabhsingh1408/banking_app/service"
)

type Customers struct {
	Name string `json:"first_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Pin  string `json:"pin" xml:"pin"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
