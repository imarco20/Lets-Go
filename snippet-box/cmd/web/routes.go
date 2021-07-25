package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", app.session.Enable(app.nosurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createSnippetForm))))))
	mux.Post("/snippet/create", app.session.Enable(app.nosurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.createSnippet))))))
	mux.Get("/snippet/:id", app.session.Enable(app.nosurf(app.authenticate(http.HandlerFunc(app.showSnippet)))))

	mux.Get("/user/signup", app.session.Enable(app.nosurf(app.authenticate(http.HandlerFunc(app.signupUserForm)))))
	mux.Post("/user/signup", app.session.Enable(app.nosurf(app.authenticate(http.HandlerFunc(app.signupUser)))))
	mux.Get("/user/login", app.session.Enable(app.nosurf(app.authenticate(http.HandlerFunc(app.loginUserForm)))))
	mux.Post("/user/login", app.session.Enable(app.nosurf(app.authenticate(http.HandlerFunc(app.loginUser)))))
	mux.Post("/user/logout", app.session.Enable(app.nosurf(app.authenticate(app.requireAuthenticatedUser(http.HandlerFunc(app.logoutUser))))))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// wrap the servermux with the secureHeaders middleware
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
