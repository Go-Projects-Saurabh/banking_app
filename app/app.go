package app

import "net/http"

func Start() {
	mux := http.NewServeMux()
	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/cutomers", getCustomers)
	//starting server
	http.ListenAndServe("localhost:5005", mux)
}
