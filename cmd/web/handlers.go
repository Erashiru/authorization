package main

import (
	"html/template"
	"net/http"

	"authorization/internal/validator"
)

type userSignupForm struct {
	Name     string
	Email    string
	Password string
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.serveError(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		app.serveError(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	t, err := template.ParseFiles("/ui/templates/index.html")
	if err != nil {
		app.serveError(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	t.Execute(w, nil)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/user/signup" {
			app.serveError(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		t, err := template.ParseFiles("/ui/templates/signup.html")
		if err != nil {
			app.serveError(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		t.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		if r.URL.Path != "/user/signup" {
			app.serveError(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		t, err := template.ParseFiles("/ui/templates/signup.html")
		if err != nil {
			app.serveError(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		t.Execute(w, nil)
	} else {
		app.serveError(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
}
