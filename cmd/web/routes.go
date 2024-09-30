package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /note/view/{id}", app.noteView)
	mux.HandleFunc("GET /note/create", app.noteCreate)
	mux.HandleFunc("POST /note/create", app.noteCreatePost)

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
