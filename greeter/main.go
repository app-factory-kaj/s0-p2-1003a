package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var g Greeting
	if name == "" {
		g = Greeting{Message: "Hello, World!"}
	} else {
		g = Greeting{Name: name, Message: "Hello, " + name + "!"}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(g); err != nil {
		log.Printf("encode greeting: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello)

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
