package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customers struct {
	Name string `json:"first_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Pin  string `json:"pin" xml:"pin"`
}

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getcutomers", getCustomers)
	http.HandleFunc("/cutomers", getCustomersXML)
	//starting server
	http.ListenAndServe("localhost:5005", nil)
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

func getCustomersXML(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{Name: "Saurabh", City: "Noida", Pin: "201301"},
		{Name: "Kavyansh", City: "Noida", Pin: "201301"},
	}
	w.Header().Add("Content-Type", "xml")
	// json.NewEncoder(w).Encode(customers)
	xml.NewEncoder(w).Encode(customers)
}
