package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /note/view/{id}", noteView)
	mux.HandleFunc("GET /note/create", noteCreate)
	mux.HandleFunc("POST /note/create", noteCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
