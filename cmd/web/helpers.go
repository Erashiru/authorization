package main

import "net/http"

func (app *application) serveError(w http.ResponseWriter, r *http.Request, errorCode int, errorMsg string) {
	w.Write([]byte(errorMsg))
	w.Write([]byte(string(errorCode)))
}
