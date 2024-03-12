package main

import (
	"authorization/internal/models"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	users *models.UserModel
}

var dsn = "./storage/storage.db"

func main() {
	db, err := models.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	app := &application{
		users: &models.UserModel{DB: db},
	}

	srv := &http.Server{
		Handler: app.routes(),
	}

	fmt.Println("Starting server on http://localhost:6666/")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
