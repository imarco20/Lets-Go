package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.showSnippet)))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// wrap the servermux with the secureHeaders middleware
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
