package main

import (
	"log"
	"net/http"

	"github.com/frankdb/urlshortener/internal/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the URL Shortener Service!"))
	})
	http.HandleFunc("/shorten", api.ShortenHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
