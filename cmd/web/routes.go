package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/user/login", app.login)
	mux.HandleFunc("/user/signup", app.signup)

	return mux
}
