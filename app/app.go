package app

import "net/http"

func Start() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/cutomers", getCustomers)
	//starting server
	http.ListenAndServe("localhost:5005", nil)
}
