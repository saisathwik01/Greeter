package main

import (
	"encoding/json"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := GreetingResponse{Message: "Hello, " + name + "!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func server() {
	http.HandleFunc("/greet", greetHandler)
	http.ListenAndServe(":8080", nil) // Start server on port 8080
}
