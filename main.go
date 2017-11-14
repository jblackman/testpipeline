package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
)

func main() {
	address := defEnv("ADDRESS", "0.0.0.0")
	port := defEnv("PORT", "7516")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received: %s %s", r.Method, r.URL.Path)
		fmt.Fprintf(w, "<html><body><h1>Hello %s!</h1></body></html>", r.URL.Path)
	})

	log.Printf("Listening on %s:%s", address, port)
	log.Fatal(http.ListenAndServe(address + ":" + port, nil))
}

func defEnv(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
