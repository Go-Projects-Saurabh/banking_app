package app

import (
	"encoding/json"
	"encoding/xml"
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
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
