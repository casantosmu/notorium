package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
	mux.Handle("GET /note/view/{id}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.noteView)))
	mux.Handle("GET /note/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.noteCreate)))
	mux.Handle("POST /note/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.noteCreatePost)))

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
