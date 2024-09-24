package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Notorium!"))
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific note with ID %d...", id)
	w.Write([]byte(msg))
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new note..."))
}

func noteCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new note..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /notes/view/{id}", noteView)
	mux.HandleFunc("GET /notes/create", noteCreate)
	mux.HandleFunc("POST /notes/create", noteCreatePost)

	fmt.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
