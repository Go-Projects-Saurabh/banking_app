package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	Name string `json:"first_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Pin  string `json:"pin" xml:"pin"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{Name: "Saurabh", City: "Noida", Pin: "201301"},
		{Name: "Kavyansh", City: "Noida", Pin: "201301"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create customer received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}
