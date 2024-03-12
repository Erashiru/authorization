package models

import "database/sql"

func NewDB(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE CONSTRAINT users_uc_username,
		email TEXT NOT NULL UNIQUE CONSTRAINT users_uc_email
		hashed_password TEXT NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
